package messaging

import "strings"

const (
	ChannelSize        = 16
	FragmentIdentifier = "#"
)

// NewAgent - agent constructor
type NewAgent func() Agent

// Agent - agent
type Agent interface {
	Name() string
	Message(m *Message)
}

func Shutdown(agent Agent) {
	if agent != nil {
		agent.Message(ShutdownMessage)
	}
}

func Startup(agent Agent) {
	if agent != nil {
		agent.Message(StartupMessage)
	}
}

func Name(agent Agent) string {
	if agent == nil {
		return ""
	}
	return getName(agent.Name())
}

func getName(uri string) string {
	i := strings.Index(uri, FragmentIdentifier)
	if i == -1 {
		return uri
	}
	return uri[:i]
}
