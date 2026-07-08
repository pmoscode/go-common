package command

// WithCommand sets the command (executable) to run.
func WithCommand(command string) func(parameter *Parameter) {
	return func(e *Parameter) {
		e.command = command
	}
}

// WithParam adds an unmasked entry (e.g. a flag like "--verbose" or a plain value)
// to the command. WithParam and WithValue are intentional aliases of each other –
// both add an unmasked entry. Use whichever name reads better at the call site
// (e.g. WithParam("--verbose") for flags, WithValue("hello") for values).
func WithParam(parameter string) func(parameter *Parameter) {
	return withEntry(parameter, false)
}

// WithValue adds an unmasked entry (e.g. a path or argument value) to the command.
// See [WithParam] for details – both are aliases of each other.
func WithValue(parameter string) func(parameter *Parameter) {
	return withEntry(parameter, false)
}

// WithValueMasked adds a value entry to the command that will be masked (***) in log output.
func WithValueMasked(parameter string) func(parameter *Parameter) {
	return withEntry(parameter, true)
}

func withEntry(entry string, masked bool) func(parameter *Parameter) {
	return func(e *Parameter) {
		e.args = append(e.args, argument{value: entry, masked: masked})
	}
}
