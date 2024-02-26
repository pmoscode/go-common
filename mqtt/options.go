package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Option func(options *mqtt.ClientOptions)

func WithBroker(ip string, port int) Option {
	return func(options *mqtt.ClientOptions) {
		options.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
	}
}

func WithClientId(clientId string) Option {
	return func(options *mqtt.ClientOptions) {
		options.SetClientID(clientId)
	}
}

func WithOrderMatters(orderMatters bool) Option {
	return func(options *mqtt.ClientOptions) {
		options.SetOrderMatters(orderMatters)
	}
}
