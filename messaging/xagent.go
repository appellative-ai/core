package messaging

import "fmt"

const (
	registerEvent    = "core:event/agent/register"
	contentTypeAgent = "application/x-agent"
	nameFmt          = "core:agent/exchange%v"
)

type ExchangeAgent interface {
	Agent
}

type agentT struct {
	name string
	ex   *Exchange
}

func NewExchangeAgent(nss string) ExchangeAgent {
	a := new(agentT)
	a.name = fmt.Sprintf(nameFmt, nss)
	a.ex = NewExchange()
	return a
}

func (a *agentT) Uri() string {
	return a.name
}

func (a *agentT) String() string { return a.Uri() }

func (a *agentT) Message(m *Message) {
	switch m.Name() {
	case ShutdownEvent:
		a.ex.Broadcast(m)
	case StartupEvent:
		a.ex.Broadcast(m)
	case registerEvent:
		agent := RegisterAgentContent(m)
		if agent != nil {
			a.ex.Register(agent)
		}
	default:
		a.ex.Message(m)
		//fmt.Printf("exchange agent - invalid name %v\n", m)
	}
}

func NewRegisterAgentMessage(a Agent) *Message {
	m := NewMessage(ChannelControl, registerEvent)
	m.SetContent(contentTypeAgent, a)
	return m
}

func RegisterAgentContent(m *Message) Agent {
	if m.Name() != registerEvent || m.ContentType() != contentTypeAgent {
		return nil
	}
	if s, ok := m.Body.(Agent); ok {
		return s
	}
	return nil
}
