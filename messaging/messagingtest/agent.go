package messagingtest

import (
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	name string
	ch   *messaging.Channel
}

func NewAgent(name string) messaging.Agent {
	a := new(agentT)
	a.name = name
	a.ch = messaging.NewEmissaryChannel()
	return a
}

func NewAgentWithChannel(name string, ch *messaging.Channel) messaging.Agent {
	a := new(agentT)
	a.name = name
	a.ch = ch
	return a
}

func (t *agentT) Name() string { return t.name }
func (t *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Name() == messaging.StartupEvent {
		t.run()
		return
	}
	t.ch.C <- m
}
func (t *agentT) run() {
	go func() {
		for {
			select {
			case msg := <-t.ch.C:
				switch msg.Name() {
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
