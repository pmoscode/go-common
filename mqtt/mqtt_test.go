package mqtt

import (
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnection(t *testing.T) {
	hostConfig := NewBrokerBuilder().
		WithHost("localhost").
		WithPort(1883).
		WithProtocol(MqttTcp).
		WithUsernameAndPassword("test", "pwd")

	brokerOpt, err := hostConfig.Build()
	require.NoError(t, err)

	client := NewClient(
		WithClientId("test"),
		brokerOpt,
	)

	mockClientImpl := &mockClient{}
	err = client.connect(mockClientImpl)
	require.NoError(t, err)

	err = client.Disconnect()
	require.NoError(t, err)

	assert.Equal(t, 1, mockClientImpl.cntConnect)
	assert.Equal(t, 1, mockClientImpl.cntDisconnect)
}

func TestPublish(t *testing.T) {
	hostConfig := NewBrokerBuilder().
		WithHost("localhost").
		WithPort(1883).
		WithProtocol(MqttTcp)

	brokerOpt, err := hostConfig.Build()
	require.NoError(t, err)

	client := NewClient(
		WithClientId("test"),
		brokerOpt,
	)

	mockClientImpl := &mockClient{}
	err = client.connect(mockClientImpl)
	require.NoError(t, err)

	client.Publish(&Message{
		Topic: "/test/testMessage",
		Value: "{'test': 2}",
	})

	err = client.Disconnect()
	require.NoError(t, err)

	assert.Equal(t, 1, mockClientImpl.cntConnect)
	assert.Equal(t, 1, mockClientImpl.cntDisconnect)
	assert.Equal(t, 1, mockClientImpl.cntPublish)
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
	cntWait int
	//lint:ignore U1000 future use possible
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
