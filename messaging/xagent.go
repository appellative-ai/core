package messaging

import "fmt"

const (
	exchangeNameFmt = "common:core:agent/exchange%v"
)

type ExchangeAgent interface {
	Agent
	Register(agent Agent)
}

type exAgentT struct {
	name string
	ex   *Exchange
}

func NewExchangeAgent(nss string) ExchangeAgent {
	a := new(exAgentT)
	a.name = fmt.Sprintf(exchangeNameFmt, nss)
	a.ex = NewExchange()
	return a
}

func (a *exAgentT) Name() string {
	return a.name
}

func (a *exAgentT) Register(agent Agent) {
	a.ex.Register(agent)
}

func (a *exAgentT) String() string { return a.Name() }

func (a *exAgentT) Message(m *Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case ShutdownEvent:
		a.ex.Broadcast(m)
	case StartupEvent:
		a.ex.Broadcast(m)
	default:
		a.ex.Message(m)
		//fmt.Printf("exchange agent - invalid name %v\n", m)
	}
}

/*
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


*/
