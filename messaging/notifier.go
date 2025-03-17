package messaging

import (
	"fmt"
	"time"
)

type NotifyItem interface {
	AgentId() string
	Type() string
	Status() string
	Message() string
}

type NotifyFunc func(e NotifyItem)

func Notify(e NotifyItem) {
	fmt.Printf("notify-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), e.AgentId(), e.Type(), e.Status(), e.Message())
}

type Notifier interface {
	Notify(status *Status)
}
