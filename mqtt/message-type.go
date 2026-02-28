package mqtt

import (
	"encoding/json"
	"fmt"
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

// ToStruct converts the received JSON message to the given target struct.
// Returns an error if the message value is not a byte slice or not valid JSON.
func (m Message) ToStruct(target any) error {
	message, err := m.ToRawString()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(message), &target); err != nil {
		return fmt.Errorf("message is not valid JSON: %w", err)
	}
	return nil
}

// ToRawString converts the received message value to an ordinary string.
// Returns an error if the value is not a byte slice.
func (m Message) ToRawString() (string, error) {
	if v, ok := m.Value.([]uint8); ok {
		return string(v), nil
	}
	return "", fmt.Errorf("message value is not a byte slice, got %T", m.Value)
}
