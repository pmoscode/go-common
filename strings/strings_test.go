package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrintJson(t *testing.T) {
	testObj := struct {
		Key1 string
		Key2 string
		Key3 string
	}{
		Key1: "val1",
		Key2: "val2",
		Key3: "val3",
	}

	result := PrettyPrintJson(testObj)
	lines := countRune(result, '\n') + 1

	assert.Equal(t, 5, lines, "JSON pretty print should produce 5 lines")
}

func TestPrettyPrintJsonContent(t *testing.T) {
	testObj := struct {
		Name string
		Age  int
	}{Name: "Test", Age: 42}

	result := PrettyPrintJson(testObj)
	assert.Contains(t, result, `"Name": "Test"`)
	assert.Contains(t, result, `"Age": 42`)
}

func TestPrettyPrintJsonEmptyStruct(t *testing.T) {
	testObj := struct{}{}
	result := PrettyPrintJson(testObj)
	assert.Equal(t, "{}", result)
}

func TestPrettyPrintYaml(t *testing.T) {
	testObj := struct {
		Key1 string
		Key2 string
		Key3 string
	}{
		Key1: "val1",
		Key2: "val2",
		Key3: "val3",
	}

	result := PrettyPrintYaml(testObj)
	lines := countRune(result, '\n')

	assert.Equal(t, 2, lines, "YAML pretty print should produce 2 newlines")
}

func TestPrettyPrintYamlContent(t *testing.T) {
	testObj := struct {
		Name string
		Port int
	}{Name: "app", Port: 3000}

	result := PrettyPrintYaml(testObj)
	assert.Contains(t, result, "name: app")
	assert.Contains(t, result, "port: 3000")
}

func TestPrettyPrintYamlEmptyStruct(t *testing.T) {
	testObj := struct{}{}
	result := PrettyPrintYaml(testObj)
	assert.Equal(t, "{}", result)
}

func TestPrettyPrintJsonSimpleTypes(t *testing.T) {
	result := PrettyPrintJson("hello")
	assert.Equal(t, `"hello"`, result)
}

func TestPrettyPrintJsonSlice(t *testing.T) {
	result := PrettyPrintJson([]int{1, 2, 3})
	assert.Contains(t, result, "1")
	assert.Contains(t, result, "2")
	assert.Contains(t, result, "3")
}

func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
