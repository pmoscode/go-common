package mqtt

import "time"

type StatusType string

const (
	Info  StatusType = "info"
	Error StatusType = "error"
	Warn  StatusType = "warn"
)

type Message struct {
	Topic string
	Value interface{}
}
type StatusPublishMessage struct {
	Type      StatusType `json:"type" binding:"required"`
	Status    string     `json:"status"`
	Timestamp time.Time  `json:"timestamp"`
}
