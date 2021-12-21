package generator

import (
	"os"
	"strings"
)

func Module(ctx *Ctx, m map[string]*FileTemplate, name, path string) error {
	ctx.Set("module_id", name)

	base := path + string(os.PathSeparator) + name

	var err error
	for key, ft := range m {
		err = ft.Create(ctx, base+string(os.PathSeparator)+key)
		if err != nil {
			return err
		}
	}

	base += string(os.PathSeparator) + "src"
	err = os.MkdirAll(base+string(os.PathSeparator)+"main"+string(os.PathSeparator)+"java", os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(base+string(os.PathSeparator)+"main"+string(os.PathSeparator)+"resources", os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(base+string(os.PathSeparator)+"test"+string(os.PathSeparator)+"java", os.ModePerm)
	if err != nil {
		return err
	}

	base += string(os.PathSeparator) + "main" + string(os.PathSeparator) + "java"
	return os.MkdirAll(base+string(os.PathSeparator)+strings.ReplaceAll(ctx.Get("group_id"), ".", string(os.PathSeparator))+string(os.PathSeparator)+name, os.ModePerm)
}

func ModulesString(ids ...string) string {
	array := make([]string, len(ids))
	for key, id := range ids {
		array[key] = "\t\t<module>" + id + "</module>"
	}

	return strings.Join(array, "\n")
}
