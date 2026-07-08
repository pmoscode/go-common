package command

import (
	"fmt"
	"strings"
)

// argument represents a single command argument, together with whether it
// should be masked (e.g. "***") when the Parameter is rendered as a string
// for logging purposes.
type argument struct {
	value  string
	masked bool
}

// Parameter holds a command and its arguments. Using a single []argument slice
// (instead of two parallel slices for values and masks) makes it impossible for
// the two to accidentally get out of sync.
type Parameter struct {
	command string
	args    []argument
}

// values returns the raw (unmasked) argument values, in order, e.g. for passing
// to exec.Command.
func (e *Parameter) values() []string {
	values := make([]string, len(e.args))
	for i, a := range e.args {
		values[i] = a.value
	}
	return values
}

func (e *Parameter) String() string {
	parts := make([]string, len(e.args))

	for idx, a := range e.args {
		if a.masked {
			parts[idx] = "***"
		} else {
			parts[idx] = a.value
		}
	}

	return fmt.Sprintf("%s %s", e.command, strings.Join(parts, " "))
}
