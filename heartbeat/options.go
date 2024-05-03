package heartbeat

type Option func(options *HeartBeat)

// WithNoWait The Heartbeat starts usually by skipping the first execution and therefore starts after the first interval time.
// With this, the execution can be done immediately.
func WithNoWait() Option {
	return func(options *HeartBeat) {
		options.noWait = true
	}
}
