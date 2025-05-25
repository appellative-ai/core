package messaging

import (
	"fmt"
	"github.com/behavioral-ai/core/fmtx"
	"time"
)

const (
	ContentTypeDispatcher = "application/x-dispatcher"
)

func NewDispatcherMessage(dispatcher Dispatcher) *Message {
	m := NewMessage(ChannelControl, ConfigEvent)
	m.SetContent(ContentTypeDispatcher, dispatcher)
	return m
}

func DispatcherContent(m *Message) (Dispatcher, bool) {
	if m.Name() != ConfigEvent || m.ContentType() != ContentTypeDispatcher {
		return nil, false
	}
	if v, ok := m.Body.(Dispatcher); ok {
		return v, true
	}
	return nil, false
}

type Dispatcher interface {
	Dispatch(agent Agent, channel any, event string)
}

func Dispatch(agent Agent, dispatcher Dispatcher, channel any, event string) {
	if dispatcher == nil || agent == nil || channel == nil {
		return
	}
	if ch, ok := channel.(*Channel); ok {
		dispatcher.Dispatch(agent, "ch:"+ch.Name(), event)
		return
	}
	if t, ok := channel.(*Ticker); ok {
		dispatcher.Dispatch(agent, "tk:"+t.Name(), event)
	}
}

type traceDispatch struct {
	allEvents bool
	channel   string
	m         map[string]string
}

func (t *traceDispatch) validEvent(event string) bool {
	if t.allEvents {
		return true
	}
	if _, ok := t.m[event]; ok {
		return true
	}
	return false
}

func (t *traceDispatch) validChannel(channel string) bool {
	if t.channel == "" {
		return true
	}
	return t.channel == channel
}

func (t *traceDispatch) Dispatch(agent Agent, channel any, event string) {
	if !t.validEvent(event) {
		return
	}
	name := "<nil>"
	if ch, ok := channel.(*Channel); ok {
		name = "ch:" + ch.Name()
		if !t.validChannel(ch.Name()) {
			return
		}
	} else {
		if tk, ok1 := channel.(*Ticker); ok1 {
			name = "tk:" + tk.Name()
		}
	}
	id := "<nil>"
	if agent != nil {
		id = agent.Uri()
	}
	fmt.Printf("trace -> %v [%v] [%v] [%v]\n", fmtx.FmtRFC3339Millis(time.Now().UTC()), id, name, event)
	//} else {
	//	fmt.Printf("trace -> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), channel, eventing, id, activity)
	//}
}

func NewTraceDispatcher() Dispatcher {
	return NewFilteredTraceDispatcher(nil, "")
}

func NewFilteredTraceDispatcher(events []string, channel string) Dispatcher {
	t := new(traceDispatch)
	if len(events) == 0 {
		t.allEvents = true
	} else {
		t.m = make(map[string]string)
		for _, event := range events {
			t.m[event] = ""
		}
	}
	t.channel = channel
	return t
}
