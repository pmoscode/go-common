package logging

import (
	"os"

	"github.com/pmoscode/go-common/process"
)

// NewLogger creates a new Logger instance with the given options.
// Options can be used to customize the logger (e.g. name, severity, writer).
func NewLogger(options ...func(*Logger)) *Logger {
	logWriter := os.Stderr

	name, err := process.GetExecutableName()
	if err != nil {
		name = "unknown"
	}

	logger := &Logger{
		name:            name,
		debug:           false,
		trace:           false,
		extend:          "",
		nameSpacing:     10,
		severitySpacing: 7,
		writer:          logWriter,
	}

	for _, o := range options {
		o(logger)
	}

	return logger
}
