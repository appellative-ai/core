package messaging

const (
	ChannelSize = 16
)

// OnShutdown - add functions to be run on shutdown
type OnShutdown interface {
	Add(func())
}

// Agent - intelligent agent
// TODO : Track agent assignment as part of the URI or separate identifier??
// //Uri() string
//
//	//Message(m *Message)
//	Track agent NID or class/type?
type Agent interface {
	Mailbox
	Finalizer
	Run()
	Shutdown()
}

func AgentCast(agent any) Agent {
	if agent == nil {
		return nil
	}
	if a, ok := agent.(Agent); ok {
		return a
	}
	return nil
}

type OpsAgent interface {
	Agent
	Notifier
	Tracer
}

func OpsAgentCast(agent any) OpsAgent {
	if agent == nil {
		return nil
	}
	if a, ok := agent.(OpsAgent); ok {
		return a
	}
	return nil
}

func AddShutdown(curr, next func()) func() {
	if next == nil {
		return nil
	}
	if curr == nil {
		curr = next
	} else {
		// !panic
		prev := curr
		curr = func() {
			prev()
			next()
		}
	}
	return curr
}

/*
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered in agent.Shutdown() : %v\n", r)
		}
	}()

*/
