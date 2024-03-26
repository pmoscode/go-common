package templates

import "text/template"

func WithFunctions(functions template.FuncMap) func(data *template.Template) *template.Template {
	return func(t *template.Template) *template.Template {
		return t.Funcs(functions)
	}
}
