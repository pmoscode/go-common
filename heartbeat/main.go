package heartbeat

import (
	"github.com/pmoscode/go-common/shutdown"
	"time"
)

type HeartBeat struct {
	interval time.Duration
	callback func()
	done     chan bool
}

func (b *HeartBeat) Run() {
	go b.beat()
}

func (b *HeartBeat) beat() {
	ticker := time.NewTicker(b.interval)
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
func New(interval time.Duration, callback func()) *HeartBeat {
	heartBeat := &HeartBeat{
		interval: interval,
		callback: callback,
		done:     make(chan bool),
	}

	shutdown.GetObserver().AddCommand(heartBeat.stop)
	shutdown.GetObserver().AddCommand(heartBeat.close)

	return heartBeat
}
