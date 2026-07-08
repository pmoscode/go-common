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

// TestParameterStringWithBrackets is a regression test for a bug where
// Parameter.String() used strings.Trim(fmt.Sprintf("%s", slice), "[]") to join
// parameter values, which incorrectly stripped literal "[" / "]" characters
// from the actual parameter values themselves.
func TestParameterStringWithBrackets(t *testing.T) {
	params := NewParameters(
		WithCommand("echo"),
		WithValue("[onlyone]"),
	)

	result := params.String()
	assert.Equal(t, "echo [onlyone]", result)
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

	assert.Equal(t, []string{"--version"}, params.values())
	assert.Equal(t, []argument{{value: "--version", masked: false}}, params.args)
}

func TestWithValue(t *testing.T) {
	params := NewParameters(
		WithCommand("echo"),
		WithValue("hello"),
	)

	assert.Equal(t, []string{"hello"}, params.values())
	assert.Equal(t, []argument{{value: "hello", masked: false}}, params.args)
}

func TestWithValueMasked(t *testing.T) {
	params := NewParameters(
		WithCommand("cmd"),
		WithValueMasked("secret"),
	)

	assert.Equal(t, []string{"secret"}, params.values())
	assert.Equal(t, []argument{{value: "secret", masked: true}}, params.args)
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
	assert.Equal(t, []string{"status"}, params.values())
}

func TestNewParametersEmpty(t *testing.T) {
	params := NewParameters()

	assert.Empty(t, params.command)
	assert.Empty(t, params.values())
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
	assert.Equal(t, []string{"--env", "production", "--token", "abc123", "--verbose"}, params.values())
	assert.Equal(t, []argument{
		{value: "--env", masked: false},
		{value: "production", masked: false},
		{value: "--token", masked: false},
		{value: "abc123", masked: true},
		{value: "--verbose", masked: false},
	}, params.args)

	str := params.String()
	assert.Contains(t, str, "***")
	assert.NotContains(t, str, "abc123")
	assert.Contains(t, str, "production")
}
