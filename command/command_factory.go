package command

import (
	"github.com/pmoscode/go-common/logging"
)

// NewCommandManager creates a new Manager with the given logger and dry-run flag.
func NewCommandManager(logger *logging.Logger, dryRun bool) *Manager {
	return &Manager{
		params: make([]func(parameter *Parameter), 0),
		logger: logger,
		dryRun: dryRun,
	}
}

// NewCommand creates a new Command with the given logger and dry-run flag.
func NewCommand(logger *logging.Logger, dryRun bool) *Command {
	return &Command{
		logger: logger,
		dryRun: dryRun,
	}
}

// NewParameters creates a new Parameter set from the given option functions.
func NewParameters(options ...func(commandOptions *Parameter)) *Parameter {
	execOptions := &Parameter{}

	for _, o := range options {
		o(execOptions)
	}

	return execOptions
}
