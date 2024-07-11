package cli

import (
	"flag"
	"os"
	"testing"
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

	if *arg.GetValue() != "Title" {
		t.Error("Expected Title, got ", *arg.GetValue())
	}

	if *file.GetValue() != "file.yaml" {
		t.Error("Expected file.yaml, got ", *file.GetValue())
	}
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

	if *arg.GetValue() != "Title" {
		t.Error("Expected Title, got ", *arg.GetValue())
	}

	if *file.GetValue() != "one.yaml" {
		t.Error("Expected one.yaml, got ", *file.GetValue())
	}
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

	if *arg.GetValue() != "Title" {
		t.Error("Expected Title, got ", *arg.GetValue())
	}

	if *file.GetValue() != "config.yaml" {
		t.Error("Expected one.yaml, got ", *file.GetValue())
	}
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

	if *arg.GetValue() != "Title" {
		t.Error("Expected Title, got ", *arg.GetValue())
	}

	if *file.GetValue() != "file.yaml" {
		t.Error("Expected file.yaml, got ", *file.GetValue())
	}

	if !*dryRun.GetValue() {
		t.Error("Expected true, got ", *dryRun.GetValue())
	}

	if *calc.GetValue() != 32.5 {
		t.Error("Expected 32.5, got ", *calc.GetValue())
	}
}
