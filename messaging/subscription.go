package messaging

import (
	"errors"
	"fmt"
)

const (
	SubscriptionCreateEvent = "event:subscription-create"
	SubscriptionCancelEvent = "event:subscription-cancel"
	ContentTypeSubscription = "application/x-subscription"
)

type Subscription struct {
	Event string
	From  string
}

type Catalog struct {
	subs []Subscription
}

func (c *Catalog) Create(s Subscription) error {
	if s.From == "" || s.Event == "" {
		return errors.New("invalid subscription: from or event is empty")
	}
	for _, item := range c.subs {
		if s.From == item.From && s.Event == item.Event {
			return errors.New(fmt.Sprintf("invalid subscription: subscription is a duplicate [%v] [%v]", s.From, s.Event))
		}
	}
	c.subs = append(c.subs, s)
	return nil
}

func (c *Catalog) CreateMessage(m *Message) error {
	if m == nil {
		return nil
	}
	if s, ok := SubscriptionCreateContent(m); ok {
		err := c.Create(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Catalog) Cancel(s Subscription) error {
	if s.From == "" || s.Event == "" {
		return errors.New("invalid subscription: from or event is empty")
	}
	for i, item := range c.subs {
		if s.From == item.From && s.Event == item.Event {
			if len(c.subs) == 1 {
				c.subs = nil
			} else {
				if i == len(c.subs)-1 {
					c.subs = c.subs[:i]
				} else {
					first := c.subs[:i]
					last := c.subs[i+1:]
					c.subs = nil
					c.subs = append(c.subs, first...)
					c.subs = append(c.subs, last...)
				}
			}
		}
	}
	return nil
}

func (c *Catalog) CancelMessage(m *Message) {
	if m == nil {
		return
	}
	if s, ok := SubscriptionCancelContent(m); ok {
		c.Cancel(s)
	}
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
