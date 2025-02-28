package messaging

const (
	ChannelSize = 16
)

// Agent - agent
type Agent interface {
	Mailbox
	Name() string
	Run()
	Shutdown()
}
