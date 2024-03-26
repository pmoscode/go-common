package templates

import (
	"fmt"
	"testing"
)

type data struct {
	One string
	Two string
}

func TestTemplateManager(t *testing.T) {
	expectedResult := "Hello Test 1! How is Test 2?"

	templateManager := NewTemplateManager[data]()
	err := templateManager.AddTemplate("test", "Hello {{ .One }}! How is {{ .Two }}?")
	if err != nil {
		t.Fatalf("Error occured %t", err)
	}

	dataObj := &data{
		One: "Test 1",
		Two: "Test 2",
	}

	result, err := templateManager.Execute("test", dataObj)
	if err != nil {
		t.Fatalf("Error occured %t", err)
	}

	if result != expectedResult {
		fmt.Println(result)
		t.Fatal("Expected: ", expectedResult, " - Got: ", result)
	}
}
