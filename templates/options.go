package templates

import "text/template"

// WithFunctions adds a [template.FuncMap] to every template in the manager.
func WithFunctions(functions template.FuncMap) func(data *template.Template) *template.Template {
	return func(t *template.Template) *template.Template {
		return t.Funcs(functions)
	}
}
