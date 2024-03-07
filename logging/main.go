package logging

import (
	"github.com/pmoscode/go-common/process"
)

func NewLogger(options ...func(*Logger)) *Logger {
	logger := &Logger{
		name:            process.GetExecutableName(),
		debug:           false,
		trace:           false,
		extend:          "",
		nameSpacing:     10,
		severitySpacing: 5,
	}

	for _, o := range options {
		o(logger)
	}

	return logger
}
