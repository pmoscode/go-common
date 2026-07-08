package logging

import (
	"fmt"
	"io"
	"sync"
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
//
// Logger is safe for concurrent use: all writes to the underlying io.Writer are
// serialized via an internal mutex.
type Logger struct {
	name            string
	nameSpacing     int
	severitySpacing int
	extend          string
	debug           bool
	trace           bool

	mu     sync.Mutex
	writer io.Writer
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

// Error logs a message with ERROR severity and appends the error (if not nil) to the
// same log line, written through the configured writer.
func (l *Logger) Error(err error, format string, params ...any) {
	if err != nil {
		l.log("ERROR", format+": %v", append(params, err)...)
		return
	}
	l.log("ERROR", format, params...)
}

// Fatal logs a message with ERROR severity, appends the error (if not nil), and exits the application.
func (l *Logger) Fatal(err error, format string, params ...any) {
	l.Error(err, format, params...)
	shutdown.Exit(1)
}

// Panic logs a message with PANIC severity and panics with the given error.
func (l *Logger) Panic(err error, format string, params ...any) {
	l.log("PANIC", format, params...)
	panic(err)
}

func (l *Logger) log(severity string, format string, params ...any) {
	entry := fmt.Sprintf(format, params...)
	dateStr := time.Now().Format("2006/01/02 15:04:05")
	line := fmt.Sprintf("%s %s%s\n", dateStr, l.header(severity), entry)

	l.writeLine(line)
}

// writeLine writes a single, already-formatted line to the underlying writer.
// It serializes access via the logger's mutex so that Logger is safe for
// concurrent use even with non-atomic writers (e.g. bytes.Buffer, files).
func (l *Logger) writeLine(line string) {
	l.mu.Lock()
	defer l.mu.Unlock()

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
