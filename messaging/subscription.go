package messaging

import (
	"errors"
	"github.com/appellative-ai/core/std"
)

const (
	SubscriptionCreateEvent = "core:event/subscription-create"
	SubscriptionCancelEvent = "core:event/subscription-cancel"
	ContentTypeSubscription = "application/x-subscription"
)

type Subscription struct {
	Path    string
	Channel string
	Name    string
	From    string
}

func NewSubscription(from, channel, name, path string) Subscription {
	return Subscription{From: from, Channel: channel, Name: name, Path: path}
}

func (s Subscription) Valid(path string) bool {
	return s.Path == "" || s.Path == path
}

type Catalog struct {
	subs []Subscription
}

func (c *Catalog) Lookup(name string) (subs []Subscription, ok bool) {
	for _, item := range c.subs {
		if name == item.Name {
			subs = append(subs, item)
			ok = true
		}
	}
	return
}

func (c *Catalog) Create(s Subscription) error {
	if s.From == "" || s.Name == "" || s.Channel == "" {
		return errors.New("invalid subscription: from or event is empty")
	}
	for _, item := range c.subs {
		// Check if already subscribed
		if s.From == item.From && s.Name == item.Name {
			return nil
		}
	}
	c.subs = append(c.subs, s)
	return nil
}

func (c *Catalog) CreateWithMessage(m *Message) error {
	if m == nil {
		return nil
	}
	if s, status := SubscriptionCreateContent(m); status.OK() {
		err := c.Create(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Catalog) Cancel(s Subscription) {
	for i, item := range c.subs {
		if s.From == item.From && s.Name == item.Name {
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
			return
		}
	}
}

func (c *Catalog) CancelWithMessage(m *Message) {
	if m == nil {
		return
	}
	if s, status := SubscriptionCancelContent(m); status.OK() {
		c.Cancel(s)
	}
}

func NewSubscriptionCreateMessage(to string, s Subscription) *Message {
	if to == "" || s.From == "" || s.Name == "" {
		return nil
	}
	// Send to publishers control channel
	m := NewMessage(ChannelControl, SubscriptionCreateEvent)
	m.AddTo(to)
	m.SetFrom(s.From)
	// Allow subscriber to determine receive channel
	if s.Channel == "" {
		s.Channel = ChannelControl
	}
	m.SetContent(ContentTypeSubscription, s)
	return m
}

func SubscriptionCreateContent(m *Message) (Subscription, *std.Status) {
	if !ValidContent(m, SubscriptionCreateEvent, ContentTypeSubscription) {
		return Subscription{}, std.NewStatus(std.StatusInvalidContent, "", nil)
	}
	return std.New[Subscription](m.Content)
}

func NewSubscriptionCancelMessage(to, from, Name string) *Message {
	if to == "" || from == "" || Name == "" {
		return nil
	}
	m := NewMessage(ChannelControl, SubscriptionCancelEvent)
	m.AddTo(to)
	m.SetFrom(from)
	m.SetContent(ContentTypeSubscription, Subscription{From: from, Name: Name})
	return m
}

func SubscriptionCancelContent(m *Message) (Subscription, *std.Status) {
	if !ValidContent(m, SubscriptionCancelEvent, ContentTypeSubscription) {
		return Subscription{}, std.NewStatus(std.StatusInvalidContent, "", nil)
	}
	return std.New[Subscription](m.Content)
}
