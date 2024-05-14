package logging

import (
	"github.com/pmoscode/go-common/process"
	"os"
)

func NewLogger(options ...func(*Logger)) *Logger {
	logWriter := os.Stderr
	logger := &Logger{
		name:            process.GetExecutableName(),
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
