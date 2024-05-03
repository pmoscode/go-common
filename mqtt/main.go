// Package mqtt provides a simple MQTT client wrapper around the paho.mqtt lib.
package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Deprecated: Use "NewClient" instead of this. Will be removed in future version.
func CreateClient(ip string, port int, clientId string) *Client {
	return NewClient(WithBroker(ip, port), WithClientId(clientId))
}

// NewClient Creates a new MQTT client with default configuration.
// To override the configuration, use the predefined options ("withBroker", "withClientId" and "withOrderMatters")
// or customize them by providing the Option type as parameter.
func NewClient(options ...Option) *Client {
	client := &Client{
		options: mqtt.NewClientOptions(),
	}

	for _, opt := range options {
		opt(client.options)
	}

	return client
}
