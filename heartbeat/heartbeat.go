// Package heartbeat implements an interval timer, which executes a function after a tick - like a heartbeat.
package heartbeat

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pmoscode/go-common/shutdown"
)

// HeartBeat Holds the function, timer and config.
type HeartBeat struct {
	interval time.Duration
	callback func()
	noWait   bool

	mu     sync.Mutex
	cancel context.CancelFunc
}

// RunForever Starts the timer and blocks the current thread until the context is cancelled
// or Stop() is called.
func (b *HeartBeat) RunForever(ctx context.Context) {
	ctx = b.setupContext(ctx)
	b.beat(ctx)
}

// Run Starts the timer without blocking the current thread.
// The heartbeat stops when the given context is cancelled or Stop() is called.
func (b *HeartBeat) Run(ctx context.Context) {
	ctx = b.setupContext(ctx)
	go b.beat(ctx)
}

// setupContext derives a cancellable context and stores the cancel function so that
// Stop() can terminate a running heartbeat regardless of whether it was started via
// Run or RunForever. Safe for concurrent use.
func (b *HeartBeat) setupContext(ctx context.Context) context.Context {
	derived, cancel := context.WithCancel(ctx)

	b.mu.Lock()
	b.cancel = cancel
	b.mu.Unlock()

	return derived
}

// beat Implements the execution logic.
// Behaves a little bit different, if the function should be executed immediately.
func (b *HeartBeat) beat(ctx context.Context) {
	if b.noWait {
		b.callback()
	}

	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			b.callback()
		}
	}
}

// Stop cancels the heartbeat context. Safe to call multiple times and concurrently.
func (b *HeartBeat) Stop() error {
	b.mu.Lock()
	cancel := b.cancel
	b.mu.Unlock()

	if cancel != nil {
		cancel()
	}
	return nil
}

// New heartbeat instance.
//   - interval takes the duration of the interval
//   - callback takes the function, which should be executed every interval
//   - options can be used to configure the instance
func New(interval time.Duration, callback func(), options ...Option) (*HeartBeat, error) {
	if interval <= 0 {
		return nil, fmt.Errorf("'interval' must be greater than 0")
	}

	heartBeat := &HeartBeat{
		interval: interval,
		callback: callback,
		noWait:   false,
	}

	for _, opt := range options {
		opt(heartBeat)
	}

	shutdown.GetObserver().AddCommand(heartBeat.Stop)

	return heartBeat, nil
}
