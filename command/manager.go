package command

import (
	"github.com/pmoscode/go-common/logging"
)

// Manager holds command parameters and provides a way to build and execute a command.
type Manager struct {
	params []func(parameter *Parameter)
	logger *logging.Logger
	dryRun bool
}

// AddParameter adds a parameter configuration function to the manager.
func (m *Manager) AddParameter(paramFn func(parameter *Parameter)) {
	m.params = append(m.params, paramFn)
}

// ExecuteCommand executes the command with all configured parameters.
func (m *Manager) ExecuteCommand() error {
	cmd := NewCommand(m.logger, m.dryRun)
	return cmd.Execute(NewParameters(m.params...))
}
