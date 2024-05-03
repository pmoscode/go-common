package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	hostConfig := NewBrokerBuilder().
		WithHost("localhost").
		WithPort(1883).
		WithProtocol(MqttTcp).
		WithUsernameAndPassword("test", "pwd")

	client := NewClient(
		WithClientId("test"),
		hostConfig.Build(),
	)

	mockClientImpl := &mockClient{}
	err := client.connect(mockClientImpl)
	if err != nil {
		t.Fatal(err)
	}

	err = client.Disconnect()
	if err != nil {
		t.Fatal(err)
	}

	if mockClientImpl.cntConnect != 1 {
		t.Fatal("Wrong connect count: ", mockClientImpl.cntConnect, " should be ", 1)
	}
	if mockClientImpl.cntDisconnect != 1 {
		t.Fatal("Wrong disconnect count: ", mockClientImpl.cntDisconnect, " should be ", 1)
	}
}

func TestPublish(t *testing.T) {
	hostConfig := NewBrokerBuilder().
		WithHost("localhost").
		WithPort(1883).
		WithProtocol(MqttTcp)

	client := NewClient(
		WithClientId("test"),
		hostConfig.Build(),
	)

	mockClientImpl := &mockClient{}
	err := client.connect(mockClientImpl)
	if err != nil {
		t.Fatal(err)
	}

	client.Publish(&Message{
		Topic: "/test/testMessage",
		Value: "{'test': 2}",
	})

	err = client.Disconnect()
	if err != nil {
		t.Fatal(err)
	}

	if mockClientImpl.cntConnect != 1 {
		t.Fatal("Wrong connect count: ", mockClientImpl.cntConnect, " should be ", 1)
	}
	if mockClientImpl.cntDisconnect != 1 {
		t.Fatal("Wrong disconnect count: ", mockClientImpl.cntDisconnect, " should be ", 1)
	}
	if mockClientImpl.cntPublish != 1 {
		t.Fatal("Wrong publish count: ", mockClientImpl.cntPublish, " should be ", 1)
	}
}

type mockClient struct {
	cntConnect    int
	cntDisconnect int
	cntPublish    int
}

func (c *mockClient) AddRoute(topic string, callback mqtt.MessageHandler) {}

func (c *mockClient) IsConnected() bool {
	return true
}

func (c *mockClient) IsConnectionOpen() bool {
	return true
}

func (c *mockClient) Connect() mqtt.Token {
	c.cntConnect++

	return &mockToken{}
}
func (c *mockClient) Disconnect(quiesce uint) {
	c.cntDisconnect++
}
func (c *mockClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	c.cntPublish++

	return &mockToken{}
}
func (c *mockClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return &mockToken{}
}
func (c *mockClient) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	return &mockToken{}
}
func (c *mockClient) Unsubscribe(topics ...string) mqtt.Token {
	return &mockToken{}
}
func (c *mockClient) OptionsReader() mqtt.ClientOptionsReader {
	return mqtt.ClientOptionsReader{}
}

type mockToken struct {
	cntWait  int
	cntError int
	complete chan struct{}
}

func (t *mockToken) Wait() bool {
	t.cntWait++

	return true
}

func (t *mockToken) WaitTimeout(time.Duration) bool {
	t.cntWait++

	return true
}

func (t *mockToken) Done() <-chan struct{} {
	return t.complete
}

func (t *mockToken) Error() error {
	return nil
}
