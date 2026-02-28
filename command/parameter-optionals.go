package command

// WithCommand sets the command (executable) to run.
func WithCommand(command string) func(parameter *Parameter) {
	return func(e *Parameter) {
		e.command = command
	}
}

// WithParam adds a parameter entry (e.g. a flag like "--verbose") to the command.
// TODO: Clarify semantic difference to WithValue – currently both have the same implementation.
func WithParam(parameter string) func(parameter *Parameter) {
	return withEntry(parameter, false)
}

// WithValue adds a value entry (e.g. a path or argument value) to the command.
// TODO: Clarify semantic difference to WithParam – currently both have the same implementation.
func WithValue(parameter string) func(parameter *Parameter) {
	return withEntry(parameter, false)
}

// WithValueMasked adds a value entry to the command that will be masked (***) in log output.
func WithValueMasked(parameter string) func(parameter *Parameter) {
	return withEntry(parameter, true)
}

func withEntry(entry string, masked bool) func(parameter *Parameter) {
	return func(e *Parameter) {
		e.params = append(e.params, entry)
		e.mask = append(e.mask, masked)
	}
}
