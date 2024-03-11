package heartbeat

import (
	"github.com/pmoscode/go-common/shutdown"
	"time"
)

type HeartBeat struct {
	interval time.Duration
	callback func()
	done     chan bool
	noWait   bool
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
	interval := b.interval

	if b.noWait {
		interval = 0
	}

	ticker := time.NewTimer(interval)
	defer ticker.Stop()

	for {
		select {
		case <-b.done:
			return
		case <-ticker.C:
			b.callback()
		}
	}
}

func (b *HeartBeat) stop() error {
	b.done <- true

	return nil
}

func (b *HeartBeat) close() error {
	close(b.done)

	return nil
}
func New(interval time.Duration, callback func(), options ...Option) *HeartBeat {
	heartBeat := &HeartBeat{
		interval: interval,
		callback: callback,
		done:     make(chan bool),
		noWait:   false,
	}

	for _, opt := range options {
		opt(heartBeat)
	}

	shutdown.GetObserver().AddCommand(heartBeat.stop)
	shutdown.GetObserver().AddCommand(heartBeat.close)

	return heartBeat
}
