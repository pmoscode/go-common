// Package heartbeat implements an interval timer, which executes a function after a tick - like a heartbeat.
package heartbeat

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/pmoscode/go-common/shutdown"
)

// HeartBeat Holds the function, timer and config.
type HeartBeat struct {
	interval   time.Duration
	callback   func()
	cancel     context.CancelFunc
	noWait     bool
	isFirstRun atomic.Bool
}

// RunForever Starts the timer and blocks the current thread until the context is cancelled.
func (b *HeartBeat) RunForever(ctx context.Context) {
	b.Run(ctx)
	<-ctx.Done()
}

// Run Starts the timer without blocking the current thread.
// The heartbeat stops when the given context is cancelled.
func (b *HeartBeat) Run(ctx context.Context) {
	ctx, b.cancel = context.WithCancel(ctx)
	go b.beat(ctx)
}

// beat Implements the execution logic.
// Behaves a little bit different, if the function should be executed immediately.
func (b *HeartBeat) beat(ctx context.Context) {
	if b.noWait {
		b.isFirstRun.Store(true)
		b.callback()
	}

	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if b.isFirstRun.Load() {
				b.isFirstRun.Store(false)
			}
			b.callback()
		}
	}
}

// Stop cancels the heartbeat context. Safe to call multiple times.
func (b *HeartBeat) Stop() error {
	if b.cancel != nil {
		b.cancel()
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
