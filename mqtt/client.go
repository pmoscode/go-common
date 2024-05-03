package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pmoscode/go-common/shutdown"
	"log"
	"time"
)

// Defines a default error message.
const notConnected = "Mqtt Client not connected! Call 'connect' method..."

// Client struct, which holds the client itself and the options for it.
type Client struct {
	client  *mqtt.Client
	options *mqtt.ClientOptions
}

// Connect connects to the mqtt broker.
// Return an error, if something went wrong.
func (c *Client) Connect() error {
	client := mqtt.NewClient(c.options)

	return c.connect(client)
}

// connect is an internal function, which handles the connection to the broker.
// Return an error, if something went wrong.
func (c *Client) connect(client mqtt.Client) error {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Could not connect to broker: ", token.Error())

		return token.Error()
	}

	c.client = &client
	log.Println("Mqtt connected to", c.options.Servers[0])

	return nil
}

// Disconnect does a clean disconnect from the mqtt broker.
func (c *Client) Disconnect() error {
	(*c.client).Disconnect(100)

	return nil
}

// Publish sends the give message to the mqtt broker.
// It uses QOS 2 and waits until arrival confirmation.
func (c *Client) Publish(message *Message) {
	if c.client == nil {
		log.Println(notConnected)
	} else {
		token := (*c.client).Publish(message.Topic, 2, false, message.FromJson())
		token.Wait()
	}
}

// Subscribe subscribes to a topic and execute a function on message arrival.
// Returns an error, if the client is not connected to a mqtt broker.
func (c *Client) Subscribe(topic string, fn func(message Message)) {
	if c.client == nil {
		log.Println(notConnected)
	} else {
		(*c.client).Subscribe(topic, 2, func(client mqtt.Client, msg mqtt.Message) {
			message := Message{
				Topic: msg.Topic(),
				Value: msg.Payload(),
			}
			fn(message)
		})
	}
}

// LoopForever Halts the current thread.
func (c *Client) LoopForever() {
	if c.client == nil {
		log.Println(notConnected)
	} else {
		shutdown.GetObserver().AddCommand(c.Disconnect)

		for {
			time.Sleep(10 * time.Second)
		}
	}
}
