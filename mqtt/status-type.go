package mqtt

import "time"

// StatusType can be used for a basic status message (like a "health message") with
type StatusType string

// Default StatusType constants.
const (
	Info  StatusType = "info"
	Error StatusType = "error"
	Warn  StatusType = "warn"
)

// StatusPublishMessage defines a possible message to represent a "health message"
type StatusPublishMessage struct {
	Type      StatusType `json:"type" binding:"required"`
	Status    string     `json:"status"`
	Timestamp time.Time  `json:"timestamp"`
}
