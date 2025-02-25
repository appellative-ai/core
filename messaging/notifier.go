package messaging

import (
	"fmt"
	"time"
)

type Event interface {
	AgentId() string
	Type() string
	Message() string
}

type NotifyFunc func(e Event)

func Notify(e Event) {
	fmt.Printf("notify-> %v [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), e.AgentId(), e.Type(), e.Message())
}

type Notifier interface {
	Notify(status *Status)
}
