package eventing

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

// NotifyFunc -
type NotifyFunc func(e NotifyEvent)

func OutputNotify(e NotifyEvent) {
	fmt.Printf("notify-> %v [%v] [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), e.Location(), e.Type(), e.RequestId(), e.Status(), e.Message())
}

func NewNotifyConfigMessage(fn NotifyFunc) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, NotifyConfigEvent)
	m.SetContent(ContentTypeNotifyConfig, fn)
	return m
}

func NotifyConfigContent(m *messaging.Message) NotifyFunc {
	if m.Event() != NotifyConfigEvent || m.ContentType() != ContentTypeNotifyConfig {
		return nil
	}
	if v, ok := m.Body.(NotifyFunc); ok {
		return v
	}
	return nil
}
