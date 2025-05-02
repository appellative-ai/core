package eventtest

import (
	"github.com/behavioral-ai/core/eventing"
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	notifier eventing.NotifyFunc
	activity eventing.ActivityFunc
}

func New() eventing.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	a.notifier = eventing.OutputNotify
	a.activity = eventing.OutputActivity
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return eventing.NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {}

func (a *agentT) AddActivity(e eventing.ActivityEvent) { a.activity(e) }

func (a *agentT) Notify(e eventing.NotifyEvent) { a.notifier(e) }
