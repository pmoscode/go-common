package shutdown

import (
	"errors"
	"testing"
)

func TestNilShutdown(t *testing.T) {
	finalize("Should not fail")
}

func TestSimpleShutdown(t *testing.T) {
	count := 0

	func1 := func() error {
		count++
		return nil
	}

	func2 := func() error {
		count = count + 2
		return nil
	}

	func3 := func() error {
		count = count + 3
		return errors.New("fake error in func3")
	}

	GetObserver().AddCommand(func1)
	GetObserver().AddCommand(func2)
	GetObserver().AddCommand(func3)

	success, err := observerSingleton.executeCommands()

	if count != 6 {
		t.Fatal("Wrong count after execution of commands: ", count, " should be ", 6)
	}

	if success != 2 {
		t.Fatal("Success execution should be: 2 not ", success)
	}

	if err != 1 {
		t.Fatal("Failed execution should be: 1 not ", err)
	}
}
