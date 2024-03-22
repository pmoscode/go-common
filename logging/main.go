package logging

import (
	"github.com/pmoscode/go-common/process"
	"log"
)

func NewLogger(options ...func(*Logger)) *Logger {
	logger := &Logger{
		name:            process.GetExecutableName(),
		debug:           false,
		trace:           false,
		extend:          "",
		nameSpacing:     10,
		severitySpacing: 7,
		writer:          log.Writer(),
	}

	for _, o := range options {
		o(logger)
	}

	return logger
}
