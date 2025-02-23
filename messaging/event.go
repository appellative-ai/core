package messaging

type Event interface {
	Name() string
	Content() string
	Source() string
	AgentId() string
}
