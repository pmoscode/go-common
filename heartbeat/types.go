package heartbeat

type PublishMessage struct {
	Alive bool `json:"alive"`
}
