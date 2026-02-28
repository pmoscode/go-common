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

func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
