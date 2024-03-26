// Package templates provides a manager, which can handle different templates of one data structure.
// So, it is ensured, that every template gets the same data structure (aka struct).
package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

// Manager holds the templates and the functions, which are applied on every template.
type Manager[T any] struct {
	templates map[string]*template.Template
	options   []func(t *template.Template) *template.Template
}

// NewTemplateManager creates a new manager with a defined struct.
// Additionally options can be applied. (Like additional template functions - ex. https://github.com/Masterminds/sprig)
func NewTemplateManager[T any](options ...func(t *template.Template) *template.Template) *Manager[T] {
	return &Manager[T]{
		templates: make(map[string]*template.Template),
		options:   options,
	}
}

// AddTemplate adds a named template with a given template string.
// If a template with a given name already exists, it won't be overwritten.
// If you wish to overwrite the template use [Manager.OverwriteTemplate]
func (m *Manager[T]) AddTemplate(templateName string, templateStr string) error {
	_, ok := m.templates[templateName]
	if !ok {
		err := m.OverwriteTemplate(templateName, templateStr)
		if err != nil {
			return err
		}
	}

	return nil
}

// OverwriteTemplate adds a named template with a given template string.
// If a template with a given name already exists, it will be overwritten!
// If you don't wish to overwrite the template use [Manager.AddTemplate]
func (m *Manager[T]) OverwriteTemplate(templateName string, templateStr string) error {
	var err error

	tmpl := template.New(templateName)

	for _, option := range m.options {
		tmpl = option(tmpl)
	}

	tmpl, err = tmpl.Parse(templateStr)
	if err != nil {
		return err
	}

	m.templates[templateName] = tmpl

	return nil
}

// Execute renders the template with the given name and data.
// It returns the rendered string or an error.
func (m *Manager[T]) Execute(templateName string, data *T) (string, error) {
	templateObj, ok := m.templates[templateName]
	if !ok {
		return "", fmt.Errorf("template '%s' does not exist", templateName)
	}

	result := bytes.NewBufferString("")
	err := templateObj.Execute(result, data)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
