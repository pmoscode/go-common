package command

import (
	"fmt"
	"strings"
)

type Parameter struct {
	command string
	params  []string
	mask    []bool
}

func (e *Parameter) String() string {
	paramsMasked := make([]string, len(e.params))

	for idx, param := range e.params {
		if e.mask[idx] {
			paramsMasked[idx] = "***"
		} else {
			paramsMasked[idx] = param
		}
	}

	params := strings.Trim(fmt.Sprintf("%s", paramsMasked), "[]")

	return fmt.Sprintf("%s %s", e.command, params)
}
