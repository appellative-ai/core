package messaging

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"time"
)

type Dispatcher interface {
	Dispatch(agent Agent, channel, event, activity string)
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

func (t *traceDispatch) Dispatch(agent Agent, channel, event, activity string) {
	if !t.validEvent(event) || !t.validChannel(channel) {
		return
	}
	id := "<nil>"
	if agent != nil {
		id = agent.Uri()
	}
	if activity == "" {
		fmt.Printf("trace -> %v [%v] [%v] [%v]\n", aspect.FmtRFC3339Millis(time.Now().UTC()), channel, event, id)
	} else {
		fmt.Printf("trace -> %v [%v] [%v] [%v] [%v]\n", aspect.FmtRFC3339Millis(time.Now().UTC()), channel, event, id, activity)
	}
}

func NewTraceDispatcher(events []string, channel string) Dispatcher {
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
