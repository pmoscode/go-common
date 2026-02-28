package templates

import (
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type data struct {
	One string
	Two string
}

func TestTemplateManager(t *testing.T) {
	expectedResult := "Hello Test 1! How is Test 2?"

	templateManager := NewTemplateManager[data]()
	err := templateManager.AddTemplate("test", "Hello {{ .One }}! How is {{ .Two }}?")
	require.NoError(t, err)

	dataObj := &data{
		One: "Test 1",
		Two: "Test 2",
	}

	result, err := templateManager.Execute("test", dataObj)
	require.NoError(t, err)

	assert.Equal(t, expectedResult, result)
}

func TestExecuteTemplateNotFound(t *testing.T) {
	mgr := NewTemplateManager[data]()
	_, err := mgr.Execute("nonexistent", &data{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func TestAddTemplateDuplicate(t *testing.T) {
	mgr := NewTemplateManager[data]()

	err := mgr.AddTemplate("dup", "first: {{ .One }}")
	require.NoError(t, err)

	// Adding with same name should NOT overwrite
	err = mgr.AddTemplate("dup", "second: {{ .Two }}")
	require.NoError(t, err)

	result, err := mgr.Execute("dup", &data{One: "A", Two: "B"})
	require.NoError(t, err)
	assert.Equal(t, "first: A", result, "original template should be preserved")
}

func TestOverwriteTemplate(t *testing.T) {
	mgr := NewTemplateManager[data]()

	err := mgr.AddTemplate("over", "first: {{ .One }}")
	require.NoError(t, err)

	err = mgr.OverwriteTemplate("over", "second: {{ .Two }}")
	require.NoError(t, err)

	result, err := mgr.Execute("over", &data{One: "A", Two: "B"})
	require.NoError(t, err)
	assert.Equal(t, "second: B", result, "template should be overwritten")
}

func TestAddTemplateInvalid(t *testing.T) {
	mgr := NewTemplateManager[data]()
	err := mgr.AddTemplate("bad", "{{ .Invalid }")
	require.Error(t, err)
}

func TestOverwriteTemplateInvalid(t *testing.T) {
	mgr := NewTemplateManager[data]()
	err := mgr.OverwriteTemplate("bad", "{{ .Invalid }")
	require.Error(t, err)
}

func TestWithFunctions(t *testing.T) {
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}

	mgr := NewTemplateManager[data](WithFunctions(funcMap))
	err := mgr.AddTemplate("funcs", `{{ upper .One }}`)
	require.NoError(t, err)

	result, err := mgr.Execute("funcs", &data{One: "hello"})
	require.NoError(t, err)
	assert.Equal(t, "HELLO", result)
}

func TestMultipleTemplates(t *testing.T) {
	mgr := NewTemplateManager[data]()

	require.NoError(t, mgr.AddTemplate("greet", "Hi {{ .One }}"))
	require.NoError(t, mgr.AddTemplate("farewell", "Bye {{ .Two }}"))

	d := &data{One: "Alice", Two: "Bob"}

	r1, err := mgr.Execute("greet", d)
	require.NoError(t, err)
	assert.Equal(t, "Hi Alice", r1)

	r2, err := mgr.Execute("farewell", d)
	require.NoError(t, err)
	assert.Equal(t, "Bye Bob", r2)
}
