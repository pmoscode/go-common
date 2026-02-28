package command

import (
	"strings"
	"testing"

	"github.com/pmoscode/go-common/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestLogger() *logging.Logger {
	return logging.NewLogger(
		logging.WithName("command-test"),
		logging.WithLogWriter(&strings.Builder{}),
	)
}

func newDebugTestLogger() *logging.Logger {
	return logging.NewLogger(
		logging.WithName("command-test"),
		logging.WithLogWriter(&strings.Builder{}),
		logging.WithSeverity(logging.DEBUG),
	)
}

// --- Parameter tests ---

func TestParameterString(t *testing.T) {
	params := NewParameters(
		WithCommand("echo"),
		WithParam("hello"),
		WithValue("world"),
	)

	result := params.String()
	assert.Equal(t, "echo hello world", result)
}

func TestParameterStringMasked(t *testing.T) {
	params := NewParameters(
		WithCommand("login"),
		WithParam("--user"),
		WithValue("admin"),
		WithParam("--password"),
		WithValueMasked("s3cret"),
	)

	result := params.String()
	assert.Contains(t, result, "login")
	assert.Contains(t, result, "--user")
	assert.Contains(t, result, "admin")
	assert.Contains(t, result, "--password")
	assert.Contains(t, result, "***")
	assert.NotContains(t, result, "s3cret")
}

func TestParameterStringEmpty(t *testing.T) {
	params := NewParameters(
		WithCommand("ls"),
	)

	result := params.String()
	assert.Equal(t, "ls ", result)
}

func TestWithCommand(t *testing.T) {
	params := NewParameters(WithCommand("git"))

	assert.Equal(t, "git", params.command)
}

func TestWithParam(t *testing.T) {
	params := NewParameters(
		WithCommand("git"),
		WithParam("--version"),
	)

	assert.Equal(t, []string{"--version"}, params.params)
	assert.Equal(t, []bool{false}, params.mask)
}

func TestWithValue(t *testing.T) {
	params := NewParameters(
		WithCommand("echo"),
		WithValue("hello"),
	)

	assert.Equal(t, []string{"hello"}, params.params)
	assert.Equal(t, []bool{false}, params.mask)
}

func TestWithValueMasked(t *testing.T) {
	params := NewParameters(
		WithCommand("cmd"),
		WithValueMasked("secret"),
	)

	assert.Equal(t, []string{"secret"}, params.params)
	assert.Equal(t, []bool{true}, params.mask)
}

// --- Command tests ---

func TestExecuteDryRun(t *testing.T) {
	logger := newTestLogger()
	cmd := NewCommand(logger, true)

	params := NewParameters(
		WithCommand("echo"),
		WithValue("hello"),
	)

	err := cmd.Execute(params)
	assert.NoError(t, err)
}

func TestExecuteSuccess(t *testing.T) {
	logger := newTestLogger()
	cmd := NewCommand(logger, false)

	params := NewParameters(
		WithCommand("echo"),
		WithValue("hello"),
	)

	err := cmd.Execute(params)
	assert.NoError(t, err)
}

func TestExecuteSuccessDebugMode(t *testing.T) {
	logger := newDebugTestLogger()
	cmd := NewCommand(logger, false)

	params := NewParameters(
		WithCommand("echo"),
		WithValue("hello"),
	)

	err := cmd.Execute(params)
	assert.NoError(t, err)
}

func TestExecuteFailure(t *testing.T) {
	logger := newTestLogger()
	cmd := NewCommand(logger, false)

	params := NewParameters(
		WithCommand("this-command-does-not-exist-at-all"),
	)

	err := cmd.Execute(params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not execute command")
}

// --- Factory tests ---

func TestNewCommand(t *testing.T) {
	logger := newTestLogger()
	cmd := NewCommand(logger, true)

	assert.NotNil(t, cmd)
	assert.True(t, cmd.dryRun)
}

func TestNewParameters(t *testing.T) {
	params := NewParameters(
		WithCommand("git"),
		WithParam("status"),
	)

	assert.Equal(t, "git", params.command)
	assert.Equal(t, []string{"status"}, params.params)
}

func TestNewParametersEmpty(t *testing.T) {
	params := NewParameters()

	assert.Empty(t, params.command)
	assert.Nil(t, params.params)
}

// --- Manager tests ---

func TestManagerAddParameter(t *testing.T) {
	logger := newTestLogger()
	mgr := NewCommandManager(logger, true)

	mgr.AddParameter(WithCommand("echo"))
	mgr.AddParameter(WithValue("hello"))

	assert.Len(t, mgr.params, 2)
}

func TestManagerExecuteDryRun(t *testing.T) {
	logger := newTestLogger()
	mgr := NewCommandManager(logger, true)

	mgr.AddParameter(WithCommand("echo"))
	mgr.AddParameter(WithValue("hello"))

	err := mgr.ExecuteCommand()
	assert.NoError(t, err)
}

func TestManagerExecuteSuccess(t *testing.T) {
	logger := newTestLogger()
	mgr := NewCommandManager(logger, false)

	mgr.AddParameter(WithCommand("echo"))
	mgr.AddParameter(WithValue("hello"))

	err := mgr.ExecuteCommand()
	assert.NoError(t, err)
}

func TestManagerExecuteFailure(t *testing.T) {
	logger := newTestLogger()
	mgr := NewCommandManager(logger, false)

	mgr.AddParameter(WithCommand("this-command-does-not-exist-at-all"))

	err := mgr.ExecuteCommand()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not execute command")
}

func TestMultipleParamsWithMasking(t *testing.T) {
	params := NewParameters(
		WithCommand("deploy"),
		WithParam("--env"),
		WithValue("production"),
		WithParam("--token"),
		WithValueMasked("abc123"),
		WithParam("--verbose"),
	)

	assert.Equal(t, "deploy", params.command)
	assert.Equal(t, []string{"--env", "production", "--token", "abc123", "--verbose"}, params.params)
	assert.Equal(t, []bool{false, false, false, true, false}, params.mask)

	str := params.String()
	assert.Contains(t, str, "***")
	assert.NotContains(t, str, "abc123")
	assert.Contains(t, str, "production")
}
