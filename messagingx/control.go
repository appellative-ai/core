package messaging

import (
	"errors"
	"github.com/behavioral-ai/core/aspect"
)

type controlAgent struct {
	running    bool
	agentId    string
	ch         chan *Message
	handler    Handler
	shutdownFn func()
}

// NewControlAgent - create an agent that only listens on a control channel, and has a default AgentRun func
func NewControlAgent(uri string, handler Handler) (Agent, error) {
	if handler == nil {
		return nil, errors.New("error: control agent message handler is nil")
	}
	return newControlAgent(uri, make(chan *Message, ChannelSize), handler), nil
	//return NewAgentWithChannels(uri, nil, nil, controlAgentRun, ctrlHandler)
}

func newControlAgent(uri string, ch chan *Message, handler Handler) *controlAgent {
	c := new(controlAgent)
	c.agentId = uri
	c.ch = ch
	c.handler = handler
	return c
}

// IsFinalized - finalized
func (c *controlAgent) IsFinalized() bool { return true }

// Uri - identity
func (c *controlAgent) Uri() string { return c.agentId }

// String - identity
func (c *controlAgent) String() string { return c.Uri() }

// Message - message an agent
func (c *controlAgent) Message(msg *Message) {
	if msg == nil {
		return
	}
	switch msg.Channel() {
	case ControlChannelType:
		if c.ch != nil {
			c.ch <- msg
		}
	default:
	}
}

// Run - run the agent
func (c *controlAgent) Run() {
	if c.running {
		return
	}
	c.running = true
	go controlAgentRun(c)
}

// Shutdown - shutdown the agent
func (c *controlAgent) Shutdown() {
	if !c.running {
		return
	}
	c.running = false
	if c.shutdownFn != nil {
		c.shutdownFn()
	}
	c.Message(NewControlMessage(c.agentId, c.agentId, ShutdownEvent))
}

// Add - add a shutdown function
func (c *controlAgent) Add(f func()) {
	c.shutdownFn = AddShutdown(c.shutdownFn, f)
}

func (c *controlAgent) shutdown() {
	close(c.ch)
}

// controlAgentRun - a simple run function that only handles control messages, and dispatches via a message handler
func controlAgentRun(c *controlAgent) {
	if c == nil || c.handler == nil {
		return
	}
	// ctrlHandler Handler
	//if h, ok := state.(Handler); ok {
	//	ctrlHandler = h
	//} else {
	//	return
	//}
	for {
		select {
		case msg, open := <-c.ch:
			if !open {
				return
			}
			switch msg.Event() {
			case ShutdownEvent:
				c.handler(NewMessageWithStatus(ControlChannelType, "", "", msg.Event(), aspect.StatusOK()))
				c.shutdown()
				return
			default:
				c.handler(msg)
			}
		default:
		}
	}
}
