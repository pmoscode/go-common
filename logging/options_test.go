package logging

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoName(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithName("tester"))
	log.Info("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "INFO    [    tester]")
	assert.Contains(t, result, "### Hello World")
}

func TestInfoNameSpacing(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithName("tester"), WithNameSpacing(20))
	log.Info("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "INFO    [              tester]")
	assert.Contains(t, result, "### Hello World")
}

func TestInfoExtend(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithExtend("--> extend <--"))
	log.Info("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "--> extend <--")
	assert.Contains(t, result, "### Hello World")
}

func TestInfoSeveritySpacing(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithSeveritySpacing(20))
	log.Info("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "INFO                 [logging.test]")
	assert.Contains(t, result, "### Hello World")
}
