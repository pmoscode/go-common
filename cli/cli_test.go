package cli

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCli(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "--arg", "Title"}

	arg := NewParameter[string]("arg", "Hello", "", "")
	file := NewParameter[string]("file", "file.yaml", "", "APP_FILE")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mgr := New()
	mgr.AddStringParameter(arg)
	mgr.AddStringParameter(file)
	mgr.Parse()

	assert.Equal(t, "Title", *arg.GetValue())
	assert.Equal(t, "file.yaml", *file.GetValue())
}

func TestAdvancedCli(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "--arg", "Title"}
	t.Setenv("APP_FILE", "one.yaml")

	arg := NewParameter[string]("arg", "Hello", "", "")
	file := NewParameter[string]("file", "file.yaml", "", "APP_FILE")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mgr := New()
	mgr.AddStringParameter(arg)
	mgr.AddStringParameter(file)
	mgr.Parse()

	assert.Equal(t, "Title", *arg.GetValue())
	assert.Equal(t, "one.yaml", *file.GetValue())
}

func TestComplexCli(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "--arg", "Title", "--file", "config.yaml"}
	t.Setenv("APP_ARG", "World")
	t.Setenv("APP_FILE", "one.yaml")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	arg := NewParameter[string]("arg", "Hello", "", "APP_ARG")
	file := NewParameter[string]("file", "file.yaml", "", "APP_FILE")

	mgr := New()
	mgr.AddStringParameter(arg)
	mgr.AddStringParameter(file)
	mgr.Parse()

	assert.Equal(t, "Title", *arg.GetValue())
	assert.Equal(t, "config.yaml", *file.GetValue())
}

func TestSimpleMixedCli(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "--arg", "Title", "--dryRun", "--calc", "32.5"}

	arg := NewParameter[string]("arg", "Hello", "", "")
	file := NewParameter[string]("file", "file.yaml", "", "")
	dryRun := NewParameter[bool]("dryRun", false, "", "")
	calc := NewParameter[float64]("calc", 15.5, "", "")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mgr := New()
	mgr.AddStringParameter(arg)
	mgr.AddStringParameter(file)
	mgr.AddBoolParameter(dryRun)
	mgr.AddFloat64Parameter(calc)
	mgr.Parse()

	assert.Equal(t, "Title", *arg.GetValue())
	assert.Equal(t, "file.yaml", *file.GetValue())
	assert.True(t, *dryRun.GetValue())
	assert.Equal(t, 32.5, *calc.GetValue())
}

func TestFloat32ParameterFromCli(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "--ratio", "3.25"}

	ratio := NewParameter[float32]("ratio", 1.5, "", "APP_RATIO")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mgr := New()
	mgr.AddFloat32Parameter(ratio)
	mgr.Parse()

	assert.Equal(t, float32(3.25), *ratio.GetValue())
}

func TestFloat32ParameterFromEnv(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app"}
	t.Setenv("APP_RATIO", "7.5")

	ratio := NewParameter[float32]("ratio", 1.5, "", "APP_RATIO")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mgr := New()
	mgr.AddFloat32Parameter(ratio)
	mgr.Parse()

	assert.Equal(t, float32(7.5), *ratio.GetValue())
}

func TestFloat32ParameterDefault(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app"}

	ratio := NewParameter[float32]("ratio", 1.5, "", "")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mgr := New()
	mgr.AddFloat32Parameter(ratio)
	mgr.Parse()

	assert.Equal(t, float32(1.5), *ratio.GetValue())
}
