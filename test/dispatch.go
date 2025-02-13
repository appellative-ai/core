package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"reflect"
)

var (
	Dispatcher = new(testDispatch)
)

type testDispatch struct{}

func (t *testDispatch) OnTick(agent any, src *messaging.Ticker) {
	fmt.Printf("OnTick()  -> %v : ticker:%v\n", DispatchName(agent), DispatchName(src))
}

func (t *testDispatch) OnMessage(agent any, msg *messaging.Message, src *messaging.Channel) {
	fmt.Printf("OnMsg()   -> %v : %v channel:%v\n", DispatchName(agent), DispatchName(msg), DispatchName(src))
}

func (t *testDispatch) OnTrace(agent any, activity any) {
	fmt.Printf("OnTrace() -> %v : %v\n", DispatchName(agent), activity)
}

func DispatchName(t any) string {
	if t == nil {
		return "<nil>"
	}
	switch ptr := t.(type) {
	case messaging.Agent:
		return ptr.Uri()
	case *messaging.Ticker:
		return ptr.Name()
	case *messaging.Channel:
		return ptr.Name()
	case *messaging.Message:
		return ptr.Event()
	case error:
		return ptr.Error()
	default:
		return fmt.Sprintf("%v", reflect.TypeOf(t))
	}
}
