package messagingtest

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

type AgentT struct {
	name  string
	ch    *messaging.Channel
	runFn func()
}

func NewAgent(name string) messaging.Agent {
	return newAgent(name, nil, nil)
}

func NewAgentOverride(name string, ch *messaging.Channel, run func()) messaging.Agent {
	return newAgent(name, ch, run)
}

func newAgent(name string, ch *messaging.Channel, run func()) *AgentT {
	a := new(AgentT)
	a.name = name
	if ch != nil {
		a.ch = ch
	} else {
		a.ch = messaging.NewChannel(messaging.ChannelControl, messaging.ChannelSize)
	}
	if run != nil {
		a.runFn = run
	} else {
		a.runFn = a.run
	}
	return a
}
func (t *AgentT) Name() string { return t.name }
func (t *AgentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Name == messaging.StartupEvent {
		t.run()
		return
	}
	t.ch.C <- m
}
func (t *AgentT) run() {
	go func() {
		for {
			select {
			case msg := <-t.ch.C:
				fmt.Printf("test: agent.Message() -> %v", msg)
				switch msg.Name {
				case messaging.ShutdownEvent:
					t.ch.Close()
					return
				default:
				}
			default:
			}
		}
	}()
}
