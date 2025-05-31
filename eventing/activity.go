package eventing

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

// ActivityFunc -
type ActivityFunc func(a ActivityEvent)

func (a ActivityEvent) IsEmpty() bool {
	return a.Agent == nil
}

func OutputActivity(a ActivityEvent) {
	uri := "<nil>"
	if a.Agent != nil {
		uri = a.Agent.Name()
	}
	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), uri, a.Event, a.Source, a.Content)
}

func NewActivityConfigMessage(fn ActivityFunc) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, ActivityConfigEvent)
	m.SetContent(ContentTypeActivityConfig, fn)
	return m
}

func ActivityConfigContent(msg *messaging.Message) ActivityFunc {
	if msg == nil || msg.ContentType() != ContentTypeActivityConfig || msg.Content == nil {
		return nil
	}
	if v, ok := msg.Content.(ActivityFunc); ok {
		return v
	}
	return nil
}
