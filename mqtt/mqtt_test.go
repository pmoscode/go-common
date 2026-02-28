package mqtt

import (
	"context"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Connection tests ---

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

func TestPublishNotConnected(t *testing.T) {
	client := NewClient(WithClientId("test"))
	// Should not panic, just log
	client.Publish(&Message{Topic: "/test", Value: "hello"})
}

func TestSubscribeNotConnected(t *testing.T) {
	client := NewClient(WithClientId("test"))
	// Should not panic, just log
	client.Subscribe("/test", func(message Message) {})
}

func TestLoopForeverNotConnected(t *testing.T) {
	client := NewClient(WithClientId("test"))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	// Should return immediately since client is nil
	client.LoopForever(ctx)
}

func TestLoopForeverCancellation(t *testing.T) {
	brokerOpt, err := NewBrokerBuilder().
		WithHost("localhost").
		WithPort(1883).
		WithProtocol(MqttTcp).
		Build()
	require.NoError(t, err)

	client := NewClient(WithClientId("test"), brokerOpt)

	mockClientImpl := &mockClient{}
	err = client.connect(mockClientImpl)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		client.LoopForever(ctx)
		close(done)
	}()

	select {
	case <-done:
		// Returned after context timeout
	case <-time.After(2 * time.Second):
		t.Fatal("LoopForever did not return after context cancellation")
	}
}

// --- Message tests ---

func TestMessageFromJson(t *testing.T) {
	msg := Message{Topic: "/test", Value: map[string]int{"a": 1}}
	result := msg.FromJson()
	assert.Contains(t, result, `"a":1`)
}

func TestMessageFromJsonString(t *testing.T) {
	msg := Message{Topic: "/test", Value: "plain string"}
	result := msg.FromJson()
	assert.Equal(t, `"plain string"`, result)
}

func TestMessageToRawString(t *testing.T) {
	msg := Message{Topic: "/test", Value: []uint8("hello world")}
	result, err := msg.ToRawString()
	require.NoError(t, err)
	assert.Equal(t, "hello world", result)
}

func TestMessageToRawStringInvalidType(t *testing.T) {
	msg := Message{Topic: "/test", Value: 12345}
	_, err := msg.ToRawString()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not a byte slice")
}

func TestMessageToStruct(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
		Val  int    `json:"val"`
	}
	msg := Message{Topic: "/test", Value: []uint8(`{"name":"test","val":42}`)}

	var p payload
	err := msg.ToStruct(&p)
	require.NoError(t, err)
	assert.Equal(t, "test", p.Name)
	assert.Equal(t, 42, p.Val)
}

func TestMessageToStructInvalidType(t *testing.T) {
	msg := Message{Topic: "/test", Value: 12345}
	var p struct{}
	err := msg.ToStruct(&p)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not a byte slice")
}

func TestMessageToStructInvalidJson(t *testing.T) {
	msg := Message{Topic: "/test", Value: []uint8(`{invalid}`)}
	var p struct{}
	err := msg.ToStruct(&p)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not valid JSON")
}

// --- Builder tests ---

func TestBrokerBuilderDefaults(t *testing.T) {
	b := NewBrokerBuilder()
	assert.Equal(t, "localhost", b.host)
	assert.Equal(t, 1880, b.port)
	assert.Equal(t, protocol(MqttTcp), b.protocol)
	assert.False(t, b.skipVerification)
	assert.False(t, b.tlsConfigured)
}

func TestBrokerBuilderProtocolWebsocket(t *testing.T) {
	b := NewBrokerBuilder().WithProtocol(Ws)
	assert.Equal(t, protocol(Ws), b.protocol)
}

func TestBrokerBuilderSkipVerification(t *testing.T) {
	b := NewBrokerBuilder().WithSkipVerification()
	assert.True(t, b.skipVerification)
}

func TestBrokerBuilderTlsEmptyPaths(t *testing.T) {
	b := NewBrokerBuilder().WithTlsCertificates("ca.pem", "", "")
	// Should not set TLS when paths are empty
	assert.False(t, b.tlsConfigured)
}

func TestBrokerBuilderProtocolAutoCorrectMqtt(t *testing.T) {
	b := NewBrokerBuilder().
		WithTlsCertificates("ca.pem", "client.crt", "client.key").
		WithProtocol(MqttTcp)
	// Should auto-correct to mqtts
	assert.Equal(t, protocol(MqttsTcp), b.protocol)
}

func TestBrokerBuilderProtocolAutoCorrectWs(t *testing.T) {
	b := NewBrokerBuilder().
		WithTlsCertificates("ca.pem", "client.crt", "client.key").
		WithProtocol(Ws)
	// Should auto-correct to wss
	assert.Equal(t, protocol(Wss), b.protocol)
}

func TestBrokerBuilderTlsBuildFailsWithBadCert(t *testing.T) {
	b := NewBrokerBuilder().
		WithTlsCertificates("/nonexistent/ca.pem", "/nonexistent/client.crt", "/nonexistent/client.key")
	_, err := b.Build()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not read CA certificate")
}

func TestCreateClientDeprecated(t *testing.T) {
	client := CreateClient("localhost", 1883, "test-id")
	assert.NotNil(t, client)
}

func TestNewClientWithOrderMatters(t *testing.T) {
	client := NewClient(WithOrderMatters(true))
	assert.NotNil(t, client)
}

// --- Mock implementations ---

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
	//nolint:unused // future use possible
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
