package messaging

import (
	"fmt"
)

type Tracer interface {
	Trace(agent Agent, channel, event, activity string)
}

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent Agent, channel, event, activity string) {
	//name := "<nil>"
	//if agent != nil {
	//	name = agent.Uri()
	//}
	if agent == nil {
		fmt.Printf("%v : %v %v %v", agent, channel, event, activity)
	} else {
		fmt.Printf("%v : %v %v %v", agent.Uri(), channel, event, activity)

	}
}

type TraceFilter struct {
	Channel string
	Event   string
}

func NewTraceFilter(channel, event string) *TraceFilter {
	f := new(TraceFilter)
	f.Channel = channel
	f.Event = event
	return f
}

func (f *TraceFilter) Access(channel, event string) bool {
	return !(f.Channel != "" && f.Channel != channel) && !(f.Event != "" && f.Event != event)
}
