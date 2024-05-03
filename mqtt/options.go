package mqtt

import (
	"fmt"
	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

// Option defines the function which overrides the default options of the mqtt client.
type Option func(options *pahoMqtt.ClientOptions)

// WithBroker defines the host and port of the mqtt broker. (default port usually is 1883)
func WithBroker(ip string, port int) Option {
	return func(options *pahoMqtt.ClientOptions) {
		options.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
	}
}

// WithClientId defines the client name which will be registered on the broker.
func WithClientId(clientId string) Option {
	return func(options *pahoMqtt.ClientOptions) {
		options.SetClientID(clientId)
	}
}

// WithOrderMatters defines if the order of publish and subscribed message is respected.
// -> it indicates that messages can be delivered asynchronously
// from the client to the application and possibly arrive out of order.
// See [pahoMqtt.ClientOptions.Order] and the "SetOrderMatters" function.
func WithOrderMatters(orderMatters bool) Option {
	return func(options *pahoMqtt.ClientOptions) {
		options.SetOrderMatters(orderMatters)
	}
}
