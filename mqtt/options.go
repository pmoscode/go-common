package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
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

// WithTlsCertificates uses TLS and certificate auth for the connection.
//   - caPemPath points to the MQTT servers root CA
//   - clientCrtPath points to the public cert of the client
//   - clientKeyPath points to the private key of the client
//   - skipVerification if "true", the client will skip validating the cert chain and the hostname
//     see: [crypto/tls.Config.skipVerification]
func WithTlsCertificates(caPemPath, clientCrtPath, clientKeyPath string, skipVerification bool) Option {
	return func(options *pahoMqtt.ClientOptions) {
		certPool := x509.NewCertPool()
		ca, err := os.ReadFile(caPemPath)
		if err != nil {
			log.Fatalln(err.Error())
		}
		certPool.AppendCertsFromPEM(ca)

		clientKeyPair, err := tls.LoadX509KeyPair(clientCrtPath, clientKeyPath)
		if err != nil {
			panic(err)
		}
		tlsConfig := &tls.Config{
			RootCAs:            certPool,
			ClientAuth:         tls.NoClientCert,
			ClientCAs:          nil,
			InsecureSkipVerify: skipVerification,
			Certificates:       []tls.Certificate{clientKeyPair},
		}

		options.SetTLSConfig(tlsConfig)
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
