package messaging

type Event interface {
	Name() string
	Message() string
	Source() string
	AgentId() string
}
