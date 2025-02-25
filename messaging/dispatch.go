package messaging

import (
	"fmt"
	"time"
)

type Dispatcher interface {
	Dispatch(agent Agent, channel, event string)
}

func Dispatch(agent Agent, dispatcher Dispatcher, channel any, event string) {
	if dispatcher == nil || agent == nil || channel == nil {
		return
	}
	if ch, ok := channel.(*Channel); ok {
		dispatcher.Dispatch(agent, "channel:"+ch.Name(), event)
		return
	}
	if t, ok := channel.(*Ticker); ok {
		dispatcher.Dispatch(agent, "ticker:"+t.Name(), event)
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

func (t *traceDispatch) Dispatch(agent Agent, channel, event string) {
	if !t.validEvent(event) || !t.validChannel(channel) {
		return
	}
	id := "<nil>"
	if agent != nil {
		id = agent.Uri()
	}
	fmt.Printf("trace -> %v [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), id, channel, event)
	//} else {
	//	fmt.Printf("trace -> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), channel, event, id, activity)
	//}
}

func NewTraceDispatcher() Dispatcher {
	return NewTraceFilteredDispatcher(nil, "")
}

func NewTraceFilteredDispatcher(events []string, channel string) Dispatcher {
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
