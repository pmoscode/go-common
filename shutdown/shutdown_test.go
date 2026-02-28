package shutdown

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilShutdown(t *testing.T) {
	defer ResetForTesting()
	finalize("Should not fail")
}

func TestSimpleShutdown(t *testing.T) {
	defer ResetForTesting()

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

	success, failed := observerSingleton.executeCommands()

	assert.Equal(t, 6, count, "count after execution of all commands")
	assert.Equal(t, 2, success, "successful executions")
	assert.Equal(t, 1, failed, "failed executions")
}
