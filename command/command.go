package command

import (
	"context"
	"fmt"
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
// It is equivalent to calling [Command.ExecuteContext] with [context.Background],
// i.e. it does not support cancellation or a timeout. Use ExecuteContext directly
// if you need to bound how long the external command may run.
//
// WARNING: This function passes the command and its parameters directly to exec.Command without any
// validation or sanitization. The caller MUST ensure that all inputs are trusted and do not originate
// from untrusted user input. Passing unsanitized input may lead to command injection vulnerabilities.
func (c *Command) Execute(options *Parameter) error {
	return c.ExecuteContext(context.Background(), options)
}

// ExecuteContext runs the command defined in the given Parameter, bound to the
// given context. If ctx is cancelled or its deadline is exceeded while the
// command is running, the underlying process is killed.
//
// WARNING: See [Command.Execute] for the same command-injection caveat regarding
// unsanitized input.
func (c *Command) ExecuteContext(ctx context.Context, options *Parameter) error {
	cmd := exec.CommandContext(ctx, options.command, options.values()...)

	if c.logger.IsDebug() {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
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
