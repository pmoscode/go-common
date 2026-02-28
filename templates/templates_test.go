package templates

import (
	"testing"

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
