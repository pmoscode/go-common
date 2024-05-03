// Package heartbeat implements an interval timer, which executes a function after a tick - like a heartbeat.
package heartbeat

import (
	"fmt"
	"github.com/pmoscode/go-common/shutdown"
	"time"
)

// HeartBeat Holds the function, timer and config.
type HeartBeat struct {
	interval   time.Duration
	callback   func()
	done       chan bool
	noWait     bool
	isFirstRun bool
}

// RunForever Starts the timer and blocks the current thread
func (b *HeartBeat) RunForever() {
	b.Run()

	for {
		time.Sleep(10 * time.Second)
	}
}

// Run Starts the timer without blocking the current thread
func (b *HeartBeat) Run() {
	go b.beat()
}

// beat Implements the execution logic.
// Behaves a little bit different, if the function should be executed immediately.
// Affects the stop function.
func (b *HeartBeat) beat() {
	if b.noWait {
		b.isFirstRun = true
		b.callback()
	}

	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-b.done:
			return
		case <-ticker.C:
			if b.isFirstRun {
				b.isFirstRun = false
			}
			b.callback()
		}
	}
}

// stop Stops the timer. Does not guarantee a final execution.
func (b *HeartBeat) stop() error {
	if !b.isFirstRun {
		b.done <- true
	}

	return nil
}

// close Closes the "done" channel (-> used by stop to end the ticker)
func (b *HeartBeat) close() error {
	close(b.done)

	return nil
}

// New heartbeat instance.
//   - interval takes the duration of the interval
//   - callback takes the function, which should be executed every interval
//   - options can be used to configure the instance
func New(interval time.Duration, callback func(), options ...Option) *HeartBeat {
	if interval <= 0 {
		fmt.Println("'interval' must be greater than 0!!")

		return nil
	}

	heartBeat := &HeartBeat{
		interval:   interval,
		callback:   callback,
		done:       make(chan bool),
		noWait:     false,
		isFirstRun: false,
	}

	for _, opt := range options {
		opt(heartBeat)
	}

	shutdown.GetObserver().AddCommand(heartBeat.stop)
	shutdown.GetObserver().AddCommand(heartBeat.close)

	return heartBeat
}
