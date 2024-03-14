package heartbeat

import (
	"fmt"
	"github.com/pmoscode/go-common/shutdown"
	"time"
)

type HeartBeat struct {
	interval   time.Duration
	callback   func()
	done       chan bool
	noWait     bool
	isFirstRun bool
}

func (b *HeartBeat) RunForever() {
	b.Run()

	for {
		time.Sleep(10 * time.Second)
	}
}
func (b *HeartBeat) Run() {
	go b.beat()
}

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

func (b *HeartBeat) stop() error {
	if !b.isFirstRun {
		b.done <- true
	}

	return nil
}

func (b *HeartBeat) close() error {
	close(b.done)

	return nil
}
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
