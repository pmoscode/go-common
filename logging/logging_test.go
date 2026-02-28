package logging

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	log := NewLogger(WithLogWriter(writer))
	log.Info("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "INFO    [logging.test]  ")
	assert.Contains(t, result, "### Hello World")
}

func TestWarningDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer))
	log.Warning("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "WARNING [logging.test]  ")
	assert.Contains(t, result, "### Hello World")
}

func TestDebugDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithSeverity(DEBUG))
	log.Debug("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "DEBUG   [logging.test]  ")
	assert.Contains(t, result, "### Hello World")
}

func TestDebugDisabled(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer))
	log.Debug("Hello %s", "World")

	assert.Empty(t, builder.String(), "debug output should be empty when debug is disabled")
}

func TestTraceDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer), WithSeverity(TRACE))
	log.Trace("Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "TRACE   [logging.test]  ")
	assert.Contains(t, result, "### Hello World")
}

func TestErrorDefault(t *testing.T) {
	builder := new(strings.Builder)
	writer := stringWriter{str: builder}

	log := NewLogger(WithLogWriter(writer))
	err := errors.New("test error")
	log.Error(err, "Hello %s", "World")

	result := builder.String()
	assert.Contains(t, result, "ERROR   [logging.test]  ")
	assert.Contains(t, result, "### Hello World")
}
