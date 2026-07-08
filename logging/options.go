package logging

import "io"

// WithName sets the logger name displayed in log output.
func WithName(name string) func(l *Logger) {
	return func(l *Logger) {
		l.name = name
	}
}

// WithExtend sets an extension string that is appended to the log header.
func WithExtend(extend string) func(l *Logger) {
	return func(l *Logger) {
		l.extend = extend
	}
}

// WithSeverity sets the minimum severity level for log output.
// TRACE enables all levels, DEBUG enables debug and above.
func WithSeverity(level Level) func(l *Logger) {
	logTrace := level == TRACE
	logDebug := level == DEBUG || logTrace

	return func(l *Logger) {
		l.debug = logDebug
		l.trace = logTrace
	}
}

// WithSeveritySpacing sets the column width for the severity field in log output.
func WithSeveritySpacing(spacing int) func(l *Logger) {
	return func(l *Logger) {
		l.severitySpacing = spacing
	}
}

// WithNameSpacing sets the column width for the name field in log output.
func WithNameSpacing(spacing int) func(l *Logger) {
	return func(l *Logger) {
		l.nameSpacing = spacing
	}
}

// WithLogWriter sets the io.Writer used for log output. Defaults to os.Stderr.
func WithLogWriter(writer io.Writer) func(l *Logger) {
	return func(l *Logger) {
		l.writer = writer
	}
}
