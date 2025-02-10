package test

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent messaging.Agent, channel, event, activity string) {
	trace(agent, channel, event, activity)
}

func trace(agent messaging.Agent, channel, event, activity string) {
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
