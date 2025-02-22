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

/*
//Finalizer
	//Notifier
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

func AgentCast(agent any) Agent {
	if agent == nil {
		return nil
	}
	if a, ok := agent.(Agent); ok {
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


	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered in agent.Shutdown() : %v\n", r)
		}
	}()

*/
