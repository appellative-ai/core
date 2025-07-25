package messaging

import (
	"fmt"
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
	if !ValidContent(m, ConfigEvent, ContentTypeDispatcher) {
		return nil, false
	}
	if v, ok := m.Content.Value.(Dispatcher); ok {
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
		dispatcher.Dispatch(agent, "ch:"+ch.Name, event)
		return
	}
	if t, ok := channel.(*Ticker); ok {
		dispatcher.Dispatch(agent, "tk:"+t.Name, event)
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
	name1 := "<nil>"
	if ch, ok := channel.(*Channel); ok {
		name1 = "ch:" + ch.Name
		if !t.validChannel(ch.Name) {
			return
		}
	} else {
		if tk, ok1 := channel.(*Ticker); ok1 {
			name1 = "tk:" + tk.Name
		}
	}
	id := "<nil>"
	if agent != nil {
		id = agent.Name()
	}
	fmt.Printf("trace -> %v [%v] [%v] [%v]\n", time.Now().UTC(), id, name1, event)
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
