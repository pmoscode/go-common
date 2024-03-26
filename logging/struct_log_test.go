package logging

import (
	"fmt"
	"strings"
	"testing"
)

func TestInfoStructYaml(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	structObj := struct {
		One string
		Two string
	}{
		One: "Hello",
		Two: "World",
	}

	log := NewLogger(WithLogWriter(writer), WithName("tester"))

	log.InfoStruct(Yaml, structObj)

	result := builder.String()
	expectedResult := "INFO    [    tester]  ### one: Hello\nINFO    [    tester]  ### two: World"

	if result != expectedResult {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}
func TestInfoStructJson(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	structObj := struct {
		One string
		Two string
	}{
		One: "Hello",
		Two: "World",
	}

	log := NewLogger(WithLogWriter(writer), WithName("tester"))

	log.InfoStruct(Json, structObj)

	result := builder.String()
	expectedResult := "INFO    [    tester]  ### {\nINFO    [    tester]  ###   \"One\": \"Hello\",\nINFO    [    tester]  ###   \"Two\": \"World\"\nINFO    [    tester]  ### }"

	if result != expectedResult {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}
