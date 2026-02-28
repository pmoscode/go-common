package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pmoscode/go-common/logging"
)

// Command wraps an external command execution with logging and dry-run support.
type Command struct {
	logger *logging.Logger
	dryRun bool
}

// Execute runs the command defined in the given Parameter.
//
// WARNING: This function passes the command and its parameters directly to exec.Command without any
// validation or sanitization. The caller MUST ensure that all inputs are trusted and do not originate
// from untrusted user input. Passing unsanitized input may lead to command injection vulnerabilities.
func (c *Command) Execute(options *Parameter) error {
	cmd := exec.Command(options.command, options.params...)

	if c.logger.IsDebug() {
		cmd.Stdout = io.MultiWriter(os.Stdout)
		cmd.Stderr = io.MultiWriter(os.Stderr)
	}

	c.logger.Info("Executing command: %s", options)

	if !c.logger.IsDebug() {
		c.logger.Info("This may take a while...")
	}

	if !c.dryRun {
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("could not execute command %s: %w", options, err)
		}
	} else {
		c.logger.Info("---> Not executing command")
	}

	if !c.logger.IsDebug() {
		c.logger.Info("... done!")
	}

	return nil
}
