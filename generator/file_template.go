package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileTemplate struct {
	Body string
}

func (ft *FileTemplate) Create(ctx *Ctx, path string) error {
	body := ft.Body
	ctx.For(func(key, value string) {
		body = strings.ReplaceAll(body, "{{"+key+"}}", value)
	})

	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if !os.IsExist(err) && err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(body), os.ModePerm)
}
