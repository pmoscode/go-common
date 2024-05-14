package logging

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type stringWriter struct {
	str *strings.Builder
}

func (w stringWriter) Write(p []byte) (n int, err error) {
	return w.str.Write(p)
}

func TestInfoDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := &stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer))

	log.Info("Hello %s", name)

	result := builder.String()
	expectedResult := "INFO    [logging.test]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}
func TestWarningDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer))

	log.Warning("Hello %s", name)

	result := builder.String()
	expectedResult := "WARNING [logging.test]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}

func TestDebugDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer), WithSeverity(DEBUG))

	log.Debug("Hello %s", name)

	result := builder.String()
	expectedResult := "DEBUG   [logging.test]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}

func TestTraceDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer), WithSeverity(TRACE))

	log.Trace("Hello %s", name)

	result := builder.String()
	expectedResult := "TRACE   [logging.test]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}

func TestErrorDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer))
	err := errors.New("test error")
	log.Error(err, "Hello %s", name)

	result := builder.String()
	expectedResult := "ERROR   [logging.test]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}
