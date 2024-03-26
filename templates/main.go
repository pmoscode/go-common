package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

type Manager[T any] struct {
	templates map[string]*template.Template
	options   []func(t *template.Template) *template.Template
}

func NewTemplateManager[T any](options ...func(t *template.Template) *template.Template) *Manager[T] {
	return &Manager[T]{
		templates: make(map[string]*template.Template),
		options:   options,
	}
}

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
