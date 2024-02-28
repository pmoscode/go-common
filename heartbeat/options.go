package heartbeat

type Option func(options *HeartBeat)

func WithNoWait() Option {
	return func(options *HeartBeat) {
		options.noWait = true
	}
}
