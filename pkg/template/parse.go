package template

import (
	"io"
	"path"
	"text/template"
)

// Parse file struct
type Parse struct {
	Path            string
	GetValueFuncMap template.FuncMap
}

// Execute parse template
func (t *Parse) Execute(wr io.Writer) error {
	tmpl := template.New(path.Base(t.Path)).Funcs(t.GetValueFuncMap)
	tmpl, err := tmpl.ParseFiles(t.Path)
	if err != nil {
		return err
	}
	err = tmpl.Execute(wr, nil)
	if err != nil {
		return err
	}
	return nil
}
