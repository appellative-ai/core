package eventing

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	NamespaceName   = "unn:behavioral-ai.github.com:resiliency:agent/collective/eventing"
	defaultDuration = time.Second * 10
)

type agentT struct {
	running  bool
	duration time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
	notifier NotifyFunc
	activity ActivityFunc
}

func newAgent() *agentT {
	a := new(agentT)
	a.duration = defaultDuration

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if !a.running {
		if m.Event() == messaging.ConfigEvent || m.Event() == NotifyConfigEvent || m.Event() == ActivityConfigEvent {
			a.configure(m)
			return
		}
		if m.Event() == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Event() == messaging.ShutdownEvent {
		a.running = false
	}
	switch m.Channel() {
	case messaging.ChannelEmissary:
		a.emissary.Send(m)
	case messaging.ChannelMaster:
		a.master.Send(m)
	case messaging.ChannelControl:
		a.emissary.Send(m)
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

// Run - run the agent
func (a *agentT) run() {
	go masterAttend(a)
	go emissaryAttend(a)
}

func (a *agentT) AddActivity(e ActivityEvent) {
	if a.activity != nil {
		a.activity(e)
	} else {
		uri := ""
		if e.Agent != nil {
			uri = e.Agent.Uri()
		}
		httpAddActivity("", uri, e.Event, e.Source, e.Content)
	}
}

func (a *agentT) Notify(e NotifyEvent) {
	if a.notifier != nil {
		a.notifier(e)
	} else {
		httpNotify(e)
	}
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case ContentTypeNotifyConfig:
		if v := NotifyConfigContent(m); v != nil {
			a.notifier = v
		}
	case ContentTypeActivityConfig:
		if v := ActivityConfigContent(m); v != nil {
			a.activity = v
		}
	case messaging.ContentTypeMap:
		cfg := messaging.ConfigMapContent(m)
		if cfg == nil {
			messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
		}
		// TODO : configure
	}
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}
