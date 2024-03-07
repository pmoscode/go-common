package logging

func WithName(name string) func(logger *Logger) {
	return func(s *Logger) {
		s.name = name
	}
}

func WithDryRun(dryRun bool) func(logger *Logger) {
	var extend string

	if dryRun {
		extend = "--> DRY RUN <--"
	}
	return func(s *Logger) {
		s.extend = extend
	}
}

func WithSeverity(level Level) func(logger *Logger) {
	logTrace := level == TRACE
	logDebug := level == DEBUG || logTrace

	return func(s *Logger) {
		s.debug = logDebug
		s.trace = logTrace
	}
}

func WithSeveritySpacing(spacing int) func(logger *Logger) {
	return func(s *Logger) {
		s.severitySpacing = spacing
	}
}

func WithNameSpacing(spacing int) func(logger *Logger) {
	return func(s *Logger) {
		s.nameSpacing = spacing
	}
}
