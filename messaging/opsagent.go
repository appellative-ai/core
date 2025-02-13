package messaging

// NewNotifierAgent - create an agent for notifications and dispatch
func NewNotifierAgent(uri string, notify Notifier, tracer Tracer) OpsAgent {
	return newNotifierAgent(uri, notify, tracer)
}

type notifierAgent struct {
	agentId string
	notify  Notifier
	tracer  Tracer
}

func newNotifierAgent(uri string, notify Notifier, tracer Tracer) *notifierAgent {
	n := new(notifierAgent)
	n.agentId = uri
	n.notify = notify
	n.tracer = tracer
	return n
}

// IsFinalized - finalized
func (c *notifierAgent) IsFinalized() bool { return true }

// Uri - identity
func (c *notifierAgent) Uri() string { return c.agentId }

// String - identity
func (c *notifierAgent) String() string { return c.Uri() }

// Message - message an agent
func (c *notifierAgent) Message(_ *Message) {}

// Run - run the agent
func (c *notifierAgent) Run() {}

// Shutdown - shutdown the agent
func (c *notifierAgent) Shutdown() {}

// Notify -
func (c *notifierAgent) Notify(err error) {
	c.notify.Notify(err)
}

// Trace -
func (c *notifierAgent) Trace(agent Agent, channel, event, activity string) {
	c.tracer.Trace(agent, channel, event, activity)
}
