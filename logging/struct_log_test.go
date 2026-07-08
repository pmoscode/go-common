package logging

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Contains(t, result, "INFO    [    tester]  ### one: Hello")
	assert.Contains(t, result, "INFO    [    tester]  ### two: World")
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
	assert.Contains(t, result, "INFO    [    tester]  ###   \"One\": \"Hello\",")
	assert.Contains(t, result, "INFO    [    tester]  ###   \"Two\": \"World\"")
}

// TestInfoStructUnknownFormat is a regression test for a bug where an unknown
// OutputFormat silently produced an empty log line instead of a diagnostic message.
func TestInfoStructUnknownFormat(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithName("tester"))
	log.InfoStruct(OutputFormat(99), struct{}{})

	result := builder.String()
	assert.Contains(t, result, "unsupported OutputFormat")
}

func TestOutputFormatString(t *testing.T) {
	assert.Equal(t, "Json", Json.String())
	assert.Equal(t, "Yaml", Yaml.String())
	assert.Equal(t, "OutputFormat(99)", OutputFormat(99).String())
}
