package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"maven-generator/generator"
	"os"
)

var moduleFiles = map[string]string{
	"module_gitignore.txt": ".gitignore",
	"module_pom.xml":       "pom.xml",
}

var projectFiles = map[string]string{
	"project_gitignore.txt": ".gitignore",
	"project_pom.xml":       "pom.xml",
}

func main() {
	var err error
	modules := make(map[string]*generator.FileTemplate)
	for file, target := range moduleFiles {
		modules[target], err = Load(file)
		if err != nil {
			panic(err)
		}
	}

	project := make(map[string]*generator.FileTemplate)
	for file, target := range projectFiles {
		project[target], err = Load(file)
		if err != nil {
			panic(err)
		}
	}

	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	ctx := generator.NewContext()
	ctx.Set("project_name", config.ProjectName)
	ctx.Set("group_id", config.GroupId)
	ctx.Set("artifact_id", config.ArtifactId)
	ctx.Set("version", config.Version)
	ctx.Set("java_version", config.JavaVersion)
	ctx.Set("modules", generator.ModulesString(config.Modules...))

	base := "./work" + string(os.PathSeparator) + config.ProjectName

	err = os.MkdirAll(base, os.ModePerm)
	if !os.IsExist(err) && err != nil {
		panic(err)
	}

	for key, ft := range project {
		err = ft.Create(ctx, base+string(os.PathSeparator)+key)
		if err != nil {
			panic(err)
		}
	}

	for _, name := range config.Modules {
		err = generator.Module(ctx, modules, name, base)
		if err != nil {
			panic(err)
		}
	}

	// Get a Buffer to Write To
	outFile, err := os.Create("./work" + string(os.PathSeparator) + config.ProjectName + ".zip")
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, base+string(os.PathSeparator), "")

	if err != nil {
		panic(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		panic(err)
	}
}

func Load(file string) (*generator.FileTemplate, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return &generator.FileTemplate{Body: string(bytes)}, nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Println(err)
	}

	if len(files) == 0 {
		f, err := w.Create(baseInZip + "placeholder.txt")
		if err != nil {
			log.Println(err)
		}

		_, err = f.Write([]byte(""))
		if err != nil {
			log.Println(err)
		}
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				log.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				log.Println(err)
			}

			_, err = f.Write(dat)
			if err != nil {
				log.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
