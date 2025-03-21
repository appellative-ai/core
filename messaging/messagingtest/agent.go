package messagingtest

import (
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	uri string
	ch  *messaging.Channel
}

func NewAgent(uri string) messaging.Agent {
	a := new(agentT)
	a.uri = uri
	a.ch = messaging.NewEmissaryChannel()
	return a
}

func NewAgentWithChannel(uri string, ch *messaging.Channel) messaging.Agent {
	a := new(agentT)
	a.uri = uri
	a.ch = ch
	return a
}

func (t *agentT) Uri() string { return t.uri }
func (t *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Event() == messaging.StartupEvent {
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
				switch msg.Event() {
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
