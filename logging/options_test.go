package logging

import (
	"fmt"
	"strings"
	"testing"
)

func TestInfoName(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer), WithName("tester"))

	log.Info("Hello %s", name)

	result := builder.String()
	expectedResult := "INFO    [    tester]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}

func TestInfoNameSpacing(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer), WithName("tester"), WithNameSpacing(20))

	log.Info("Hello %s", name)

	result := builder.String()
	expectedResult := "INFO    [              tester]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}

func TestInfoExtend(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer), WithExtend("--> extend <--"))

	log.Info("Hello %s", name)

	result := builder.String()
	expectedResult := "INFO    [logging.test] --> extend <-- ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}

func TestInfoSeveritySpacing(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	name := "World"

	log := NewLogger(WithLogWriter(writer), WithSeveritySpacing(20))

	log.Info("Hello %s", name)

	result := builder.String()
	expectedResult := "INFO                 [logging.test]  ### Hello " + name

	if strings.HasSuffix(result, expectedResult) {
		fmt.Println(result)
		t.Fatalf("Output should be: '%s' but got '%s'", expectedResult, result)
	}
}
