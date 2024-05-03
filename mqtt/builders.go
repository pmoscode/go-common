package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

// protocol defines the target connection protocol
type protocol string

const (
	MqttTcp  = "mqtt://"
	MqttsTcp = "mqtts://"
	Ws       = "ws://"
	Wss      = "wss://"
)

// Broker defines the properties ,which are needed to build the broker host and the TLS connection
type Broker struct {
	host                                    string
	port                                    int
	protocol                                protocol
	caPemPath, clientCrtPath, clientKeyPath string
	username, password                      string
	skipVerification                        bool
	tlsConfigured                           bool
}

// WithHost defines the host of the broker. Can be either a DNS or an IP.
func (b *Broker) WithHost(host string) *Broker {
	b.host = host

	return b
}

// WithPort defines the port of the broker to connect to.
func (b *Broker) WithPort(port int) *Broker {
	b.port = port

	return b
}

// WithProtocol defines the protocol to use. There are these two options:
//   - tcp (normal - mqtt:// or secured - mqtts://)
//   - websocket (normal - ws:// or secured - wss://)
//
// Provides an autocorrection, when WithTlsCertificates is set.
func (b *Broker) WithProtocol(protocol protocol) *Broker {
	b.protocol = protocol

	if b.tlsConfigured {
		b.adjustProtocolToSecure()
	}

	return b
}

// WithTlsCertificates uses TLS and certificate auth for the connection.
//   - caPemPath points to the MQTT servers root CA
//   - clientCrtPath points to the public cert of the client
//   - clientKeyPath points to the private key of the client
//
// If either clientCrtPath or clientKeyPath is an empty string, nothing will be changed.
func (b *Broker) WithTlsCertificates(caPemPath, clientCrtPath, clientKeyPath string) *Broker {
	if clientKeyPath == "" || clientCrtPath == "" {
		return b
	}

	b.caPemPath = caPemPath
	b.clientCrtPath = clientCrtPath
	b.clientKeyPath = clientKeyPath

	b.adjustProtocolToSecure()
	b.tlsConfigured = true

	return b
}

// WithSkipVerification the client will skip validating the cert chain and the hostname
//
//	see: [crypto/tls.Config.skipVerification]
func (b *Broker) WithSkipVerification() *Broker {
	b.skipVerification = true

	return b
}

// WithUsernameAndPassword defines the username and password, if the broker requires authentication.
func (b *Broker) WithUsernameAndPassword(username, password string) *Broker {
	b.username = username
	b.password = password

	return b
}

// Build should be called, when all configuration is done. It creates an Option type, which is used in NewClient function.
func (b *Broker) Build() Option {
	return func(options *pahoMqtt.ClientOptions) {
		if b.username != "" || b.password != "" {
			options.SetUsername(b.username)
			options.SetPassword(b.password)
		}

		if b.tlsConfigured {
			options.SetTLSConfig(b.buildTLSConfig())
		}

		options.AddBroker(fmt.Sprintf("%s%s:%d", b.protocol, b.host, b.port))
	}
}

// buildTLSConfig is an internal function, which set up the TLS configuration.
func (b *Broker) buildTLSConfig() *tls.Config {
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(b.caPemPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	certPool.AppendCertsFromPEM(ca)

	clientKeyPair, err := tls.LoadX509KeyPair(b.clientCrtPath, b.clientKeyPath)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		RootCAs:            certPool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: b.skipVerification,
		Certificates:       []tls.Certificate{clientKeyPair},
	}
	return tlsConfig
}

// adjustProtocolToSecure is called to ensure, that the protocol matches the possible configured TLS config.
func (b *Broker) adjustProtocolToSecure() {
	if b.protocol == MqttTcp {
		b.protocol = MqttsTcp
	}

	if b.protocol == Ws {
		b.protocol = Wss
	}
}

// NewBrokerBuilder configures the connection parameters like:
//   - Host, port, protocol
//   - TLS connection
func NewBrokerBuilder() *Broker {
	return &Broker{
		host:             "localhost",
		port:             1880,
		protocol:         MqttTcp,
		caPemPath:        "",
		clientCrtPath:    "",
		clientKeyPath:    "",
		username:         "",
		password:         "",
		skipVerification: false,
		tlsConfigured:    false,
	}
}
