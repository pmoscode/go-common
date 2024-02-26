package mqtt

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pmoscode/go-common/shutdown"
	"log"
	"time"
)

const notConnected = "Mqtt Client not connected! Call 'connect' method..."

func (m Message) FromJson() string {
	marshal, err := json.Marshal(m.Value)
	if err != nil {
		return fmt.Sprintf("%v", m.Value)
	}
	return string(marshal)
}

func (m Message) ToStruct(target interface{}) {
	message := m.ToRawString()
	err := json.Unmarshal([]byte(message), &target)
	if err != nil {
		log.Println("Message is not a valid Json: ", message)
	}
}

func (m Message) ToRawString() string {
	return string(m.Value.([]uint8))
}

type Client struct {
	client  *mqtt.Client
	options *mqtt.ClientOptions
}

func (c *Client) Connect() error {
	client := mqtt.NewClient(c.options)

	return c.connect(client)
}

func (c *Client) connect(client mqtt.Client) error {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Could not connect to broker: ", token.Error())

		return token.Error()
	}

	c.client = &client
	log.Println("Mqtt connected to", c.options.Servers[0])

	return nil
}

func (c *Client) Disconnect() error {
	client := *c.client
	client.Disconnect(100)

	return nil
}

func (c *Client) Publish(message *Message) {
	if c.client == nil {
		log.Println(notConnected)
	} else {
		client := *c.client
		token := client.Publish(message.Topic, 2, false, message.FromJson())
		token.Wait()
	}
}

func (c *Client) Subscribe(topic string, fn func(message Message)) {
	if c.client == nil {
		log.Println(notConnected)
	} else {
		client := *c.client
		client.Subscribe(topic, 2, func(client mqtt.Client, msg mqtt.Message) {
			message := Message{
				Topic: msg.Topic(),
				Value: msg.Payload(),
			}
			fn(message)
		})
	}
}

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

// Deprecated: Use "NewClient" instead of this. Will be removed in future
func CreateClient(ip string, port int, clientId string) *Client {
	return NewClient(WithBroker(ip, port), WithClientId(clientId))
}

func NewClient(options ...Option) *Client {
	client := &Client{
		options: mqtt.NewClientOptions(),
	}

	for _, opt := range options {
		opt(client.options)
	}

	return client
}
