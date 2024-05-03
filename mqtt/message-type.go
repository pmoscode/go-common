package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
)

// Message defines the default message which is received and sent by the mqtt client.
// Here are two fields:
//   - Topic: The topic to which the message should be sent
//   - Value: Any struct or type (simple string or a complex type)
type Message struct {
	Topic string
	Value any
}

// FromJson converts a JSON struct to a string.
func (m Message) FromJson() string {
	marshal, err := json.Marshal(m.Value)
	if err != nil {
		return fmt.Sprintf("%v", m.Value)
	}
	return string(marshal)
}

// ToStruct converts the received JSON message to the given target struct
func (m Message) ToStruct(target any) {
	message := m.ToRawString()
	err := json.Unmarshal([]byte(message), &target)
	if err != nil {
		log.Println("Message is not a valid Json: ", message)
	}
}

// ToRawString converts the received JSON string to an ordinary string
func (m Message) ToRawString() string {
	return string(m.Value.([]uint8))
}
