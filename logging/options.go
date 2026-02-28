package logging

import "io"

// WithName sets the logger name displayed in log output.
func WithName(name string) func(logger *Logger) {
	return func(s *Logger) {
		s.name = name
	}
}

// WithExtend sets an extension string that is appended to the log header.
func WithExtend(extend string) func(logger *Logger) {
	return func(s *Logger) {
		s.extend = extend
	}
}

// WithSeverity sets the minimum severity level for log output.
// TRACE enables all levels, DEBUG enables debug and above.
func WithSeverity(level Level) func(logger *Logger) {
	logTrace := level == TRACE
	logDebug := level == DEBUG || logTrace

	return func(s *Logger) {
		s.debug = logDebug
		s.trace = logTrace
	}
}

// WithSeveritySpacing sets the column width for the severity field in log output.
func WithSeveritySpacing(spacing int) func(logger *Logger) {
	return func(s *Logger) {
		s.severitySpacing = spacing
	}
}

// WithNameSpacing sets the column width for the name field in log output.
func WithNameSpacing(spacing int) func(logger *Logger) {
	return func(s *Logger) {
		s.nameSpacing = spacing
	}
}

// WithLogWriter sets the io.Writer used for log output. Defaults to os.Stderr.
func WithLogWriter(writer io.Writer) func(logger *Logger) {
	return func(s *Logger) {
		s.writer = writer
	}
}
