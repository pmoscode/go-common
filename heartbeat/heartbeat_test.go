package heartbeat

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pmoscode/go-common/shutdown"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHeartbeat(t *testing.T) {
	defer shutdown.ResetForTesting()

	hb, err := New(100*time.Millisecond, func() {})
	require.NoError(t, err)
	assert.NotNil(t, hb)
	assert.False(t, hb.noWait)
}

func TestNewHeartbeatWithNoWait(t *testing.T) {
	defer shutdown.ResetForTesting()

	hb, err := New(100*time.Millisecond, func() {}, WithNoWait())
	require.NoError(t, err)
	assert.True(t, hb.noWait)
}

func TestNewHeartbeatInvalidInterval(t *testing.T) {
	defer shutdown.ResetForTesting()

	hb, err := New(0, func() {})
	assert.Nil(t, hb)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestNewHeartbeatNegativeInterval(t *testing.T) {
	defer shutdown.ResetForTesting()

	hb, err := New(-5*time.Second, func() {})
	assert.Nil(t, hb)
	require.Error(t, err)
}

func TestRunExecutesCallback(t *testing.T) {
	defer shutdown.ResetForTesting()

	var count atomic.Int32
	hb, err := New(50*time.Millisecond, func() {
		count.Add(1)
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	hb.Run(ctx)

	time.Sleep(180 * time.Millisecond)
	cancel()
	time.Sleep(20 * time.Millisecond)

	assert.GreaterOrEqual(t, count.Load(), int32(2), "callback should have been called at least twice")
}

func TestRunWithNoWaitExecutesImmediately(t *testing.T) {
	defer shutdown.ResetForTesting()

	var count atomic.Int32
	hb, err := New(500*time.Millisecond, func() {
		count.Add(1)
	}, WithNoWait())
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	hb.Run(ctx)

	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(20 * time.Millisecond)

	assert.GreaterOrEqual(t, count.Load(), int32(1), "callback should have been called immediately")
}

func TestRunForeverBlocksUntilCancel(t *testing.T) {
	defer shutdown.ResetForTesting()

	var count atomic.Int32
	hb, err := New(50*time.Millisecond, func() {
		count.Add(1)
	})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		hb.RunForever(ctx)
		close(done)
	}()

	select {
	case <-done:
		// RunForever returned after context cancellation
	case <-time.After(2 * time.Second):
		t.Fatal("RunForever did not return after context cancellation")
	}

	assert.GreaterOrEqual(t, count.Load(), int32(2))
}

func TestStopIsIdempotent(t *testing.T) {
	defer shutdown.ResetForTesting()

	hb, err := New(100*time.Millisecond, func() {})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	hb.Run(ctx)
	cancel()

	// Stop should be safe to call multiple times
	assert.NoError(t, hb.Stop())
	assert.NoError(t, hb.Stop())
}

func TestStopBeforeRun(t *testing.T) {
	defer shutdown.ResetForTesting()

	hb, err := New(100*time.Millisecond, func() {})
	require.NoError(t, err)

	// Stop before Run should not panic (cancel is nil)
	assert.NoError(t, hb.Stop())
}

func TestCallbackStopsAfterCancel(t *testing.T) {
	defer shutdown.ResetForTesting()

	var count atomic.Int32
	hb, err := New(30*time.Millisecond, func() {
		count.Add(1)
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	hb.Run(ctx)

	time.Sleep(100 * time.Millisecond)
	cancel()
	countAtCancel := count.Load()

	time.Sleep(100 * time.Millisecond)
	countAfterWait := count.Load()

	// After cancel, no more than 1 additional callback should have fired (race window)
	assert.LessOrEqual(t, countAfterWait-countAtCancel, int32(1))
}
