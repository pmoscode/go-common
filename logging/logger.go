package logging

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/pmoscode/go-common/shutdown"
)

// Level defines the severity level for logging.
type Level string

const (
	// INFO is the default log level.
	INFO Level = "info"
	// DEBUG enables debug log output.
	DEBUG Level = "debug"
	// WARNING is the warning log level.
	WARNING Level = "warning"
	// TRACE enables the most verbose log output.
	TRACE Level = "trace"
)

// Logger provides structured logging with configurable severity levels, name spacing, and output writer.
type Logger struct {
	name            string
	nameSpacing     int
	severitySpacing int
	extend          string
	debug           bool
	trace           bool
	writer          io.Writer
}

// Info logs a message with INFO severity.
func (l *Logger) Info(format string, params ...any) {
	l.log("INFO", format, params...)
}

// Warning logs a message with WARNING severity.
func (l *Logger) Warning(format string, params ...any) {
	l.log("WARNING", format, params...)
}

// Debug logs a message with DEBUG severity. Only produces output if debug level is enabled.
func (l *Logger) Debug(format string, params ...any) {
	if l.debug {
		l.log("DEBUG", format, params...)
	}
}

// Trace logs a message with TRACE severity. Only produces output if trace level is enabled.
func (l *Logger) Trace(format string, params ...any) {
	if l.trace {
		l.log("TRACE", format, params...)
	}
}

// Error logs a message with ERROR severity and prints the error if not nil.
func (l *Logger) Error(err error, format string, params ...any) {
	l.log("ERROR", format, params...)
	if err != nil {
		log.Println(err)
	}
}

// Fatal logs a message with ERROR severity, prints the error if not nil, and exits the application.
func (l *Logger) Fatal(err error, format string, params ...any) {
	l.log("ERROR", format, params...)
	if err != nil {
		log.Println(err)
	}
	shutdown.Exit(1)
}

// Panic logs a message with Panic severity and panics with the given error.
func (l *Logger) Panic(err error, format string, params ...any) {
	l.log("Panic", format, params...)
	panic(err)
}

func (l *Logger) log(severity string, format string, params ...any) {
	entry := fmt.Sprintf(format, params...)
	dateStr := time.Now().Format("2006/01/02 15:04:05")
	line := fmt.Sprintf("%s %s%s\n", dateStr, l.header(severity), entry)

	_, err := l.writer.Write([]byte(line))
	if err != nil {
		return
	}
}

func (l *Logger) header(severity string) string {
	return fmt.Sprintf("%-*s [%*s] %s ### ", l.severitySpacing, severity, l.nameSpacing, l.name, l.extend)
}

// IsDebug returns true if debug level logging is enabled.
func (l *Logger) IsDebug() bool {
	return l.debug
}

// IsTrace returns true if trace level logging is enabled.
func (l *Logger) IsTrace() bool {
	return l.trace
}
