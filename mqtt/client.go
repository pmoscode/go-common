package mqtt

import (
	"context"
	"log"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pmoscode/go-common/shutdown"
)

// Defines a default error message.
const notConnected = "Mqtt Client not connected! Call 'connect' method..."

// disconnectQuiesceMs is the time (ms) given to the client to finish pending work before disconnecting.
const disconnectQuiesceMs = 100

// defaultQoS is the QoS level used for Publish/Subscribe calls.
const defaultQoS byte = 2

// Client wraps the paho MQTT client with convenience methods for connect, publish, subscribe, and disconnect.
// Client is safe for concurrent use.
type Client struct {
	mu      sync.RWMutex
	client  mqtt.Client
	options *mqtt.ClientOptions

	registerShutdownHookOnce sync.Once
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
	c.mu.Lock()
	previous := c.client
	c.mu.Unlock()

	if previous != nil && previous.IsConnected() {
		previous.Disconnect(disconnectQuiesceMs)
	}

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Could not connect to broker: ", token.Error())

		return token.Error()
	}

	c.mu.Lock()
	c.client = client
	c.mu.Unlock()

	log.Println("Mqtt connected to", c.options.Servers[0])

	return nil
}

func (c *Client) getClient() mqtt.Client {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.client
}

// Disconnect does a clean disconnect from the mqtt broker.
func (c *Client) Disconnect() error {
	client := c.getClient()
	if client == nil {
		log.Println(notConnected)
		return nil
	}

	client.Disconnect(disconnectQuiesceMs)

	return nil
}

// Publish sends the give message to the mqtt broker.
// It uses QOS 2 and waits until arrival confirmation.
func (c *Client) Publish(message *Message) error {
	client := c.getClient()
	if client == nil {
		log.Println(notConnected)
		return nil
	}

	token := client.Publish(message.Topic, defaultQoS, false, message.FromJson())
	token.Wait()
	if err := token.Error(); err != nil {
		log.Println("Could not publish message: ", err)
		return err
	}

	return nil
}

// Subscribe subscribes to a topic and execute a function on message arrival.
// Returns an error, if the client is not connected to a mqtt broker or if the
// subscription itself failed.
func (c *Client) Subscribe(topic string, fn func(message Message)) error {
	client := c.getClient()
	if client == nil {
		log.Println(notConnected)
		return nil
	}

	token := client.Subscribe(topic, defaultQoS, func(client mqtt.Client, msg mqtt.Message) {
		message := Message{
			Topic: msg.Topic(),
			Value: msg.Payload(),
		}
		fn(message)
	})
	token.Wait()
	if err := token.Error(); err != nil {
		log.Println("Could not subscribe to topic: ", err)
		return err
	}

	return nil
}

// LoopForever halts the current thread until the context is cancelled.
// On cancellation, the client will be disconnected cleanly.
//
// LoopForever registers a shutdown hook (to disconnect on process exit) only once
// per Client instance, even if LoopForever is called multiple times (e.g. as part
// of a reconnect loop), to avoid unbounded growth of the global shutdown observer.
func (c *Client) LoopForever(ctx context.Context) {
	if c.getClient() == nil {
		log.Println(notConnected)
	} else {
		c.registerShutdownHookOnce.Do(func() {
			shutdown.GetObserver().AddCommand(c.Disconnect)
		})
		<-ctx.Done()
	}
}
