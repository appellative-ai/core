package messaging

// NewNotifierAgent - create an agent for notifications and dispatch
func NewNotifierAgent(uri string, notify Notifier) OpsAgent {
	return newNotifierAgent(uri, notify)
}

type notifierAgent struct {
	agentId string
	notify  Notifier
}

func newNotifierAgent(uri string, notify Notifier) *notifierAgent {
	n := new(notifierAgent)
	n.agentId = uri
	n.notify = notify
	return n
}

// IsFinalized - finalized
func (c *notifierAgent) IsFinalized() bool { return true }

// Uri - identity
func (c *notifierAgent) Uri() string { return c.agentId }

// String - identity
func (c *notifierAgent) String() string { return c.Uri() }

// Name - class identity
func (c *notifierAgent) Name() string { return c.Uri() }

// Message - message an agent
func (c *notifierAgent) Message(_ *Message) {}

// Run - run the agent
func (c *notifierAgent) Run() {}

// Shutdown - shutdown the agent
func (c *notifierAgent) Shutdown() {}

// Notify -
func (c *notifierAgent) Notify(status *Status) {
	c.notify.Notify(status)
}
