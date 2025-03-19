package messaging

import "strings"

const (
	ChannelSize          = 16
	AssignmentIdentifier = "#"
)

// Agent - agent
type Agent interface {
	Mailbox
	Run()
}

func Shutdown(agent Agent) {
	if agent != nil {
		agent.Message(ShutdownMessage)
	}
}

func Name(agent Agent) string {
	if agent == nil {
		return ""
	}
	uri := agent.Uri()
	i := strings.Index(uri, AssignmentIdentifier)
	if i == -1 {
		return uri
	}
	return uri[:i]
}
