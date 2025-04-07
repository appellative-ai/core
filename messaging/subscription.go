package messaging

const (
	SubscriptionCreateEvent = "event:subscription-create"
	SubscriptionCancelEvent = "event:subscription-cancel"
	ContentTypeSubscription = "application/x-subscription"
)

type Subscription struct {
	Event string
	From  string
}

func NewSubscriptionCreateMessage(to, from, event string) *Message {
	m := NewMessage(Control, SubscriptionCreateEvent)
	m.SetTo(to)
	m.SetFrom(from)
	m.SetContent(ContentTypeSubscription, Subscription{From: from, Event: event})
	return m
}

func SubscriptionCreateContent(m *Message) (Subscription, bool) {
	if m.Event() != SubscriptionCreateEvent || m.ContentType() != ContentTypeSubscription {
		return Subscription{}, false
	}
	if v, ok := m.Body.(Subscription); ok {
		return v, true
	}
	return Subscription{}, false
}

func NewSubscriptionCancelMessage(to, from, event string) *Message {
	m := NewMessage(Control, SubscriptionCancelEvent)
	m.SetTo(to)
	m.SetFrom(from)
	m.SetContent(ContentTypeSubscription, Subscription{From: from, Event: event})
	return m
}

func SubscriptionCancelContent(m *Message) (Subscription, bool) {
	if m.Event() != SubscriptionCancelEvent || m.ContentType() != ContentTypeSubscription {
		return Subscription{}, false
	}
	if v, ok := m.Body.(Subscription); ok {
		return v, true
	}
	return Subscription{}, false
}
