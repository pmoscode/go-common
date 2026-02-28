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

func TestGetObserverIsSingleton(t *testing.T) {
	defer ResetForTesting()

	obs1 := GetObserver()
	obs2 := GetObserver()

	assert.Same(t, obs1, obs2, "GetObserver should return the same instance")
}

func TestResetForTesting(t *testing.T) {
	GetObserver().AddCommand(func() error { return nil })
	assert.NotNil(t, observerSingleton)

	ResetForTesting()
	assert.Nil(t, observerSingleton)

	// After reset, GetObserver creates a new instance
	obs := GetObserver()
	assert.NotNil(t, obs)
	assert.Len(t, observerSingleton.functions, 0)

	defer ResetForTesting()
}

func TestFinalizeNilSingleton(t *testing.T) {
	// Ensure finalize does not panic when singleton is nil
	defer ResetForTesting()
	ResetForTesting()
	finalize("nil singleton test")
}

func TestExecuteCommandsEmpty(t *testing.T) {
	defer ResetForTesting()

	GetObserver()
	success, failed := observerSingleton.executeCommands()
	assert.Equal(t, 0, success)
	assert.Equal(t, 0, failed)
}

func TestExecuteCommandsAllSuccess(t *testing.T) {
	defer ResetForTesting()

	GetObserver().AddCommand(func() error { return nil })
	GetObserver().AddCommand(func() error { return nil })
	GetObserver().AddCommand(func() error { return nil })

	success, failed := observerSingleton.executeCommands()
	assert.Equal(t, 3, success)
	assert.Equal(t, 0, failed)
}

func TestExecuteCommandsAllFail(t *testing.T) {
	defer ResetForTesting()

	GetObserver().AddCommand(func() error { return errors.New("err1") })
	GetObserver().AddCommand(func() error { return errors.New("err2") })

	success, failed := observerSingleton.executeCommands()
	assert.Equal(t, 0, success)
	assert.Equal(t, 2, failed)
}

func TestExitOnPanicNoPanic(t *testing.T) {
	defer ResetForTesting()
	// ExitOnPanic without panic should just return
	func() {
		defer ExitOnPanic()
	}()
}

func TestFinalizeWithCommands(t *testing.T) {
	defer ResetForTesting()

	called := false
	GetObserver().AddCommand(func() error {
		called = true
		return nil
	})

	finalize("test hooks")
	assert.True(t, called)
}
