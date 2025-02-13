package test

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messagingx"
	"reflect"
)

var (
	Dispatcher = new(testDispatch)
)

type testDispatch struct{}

func (t *testDispatch) OnTick(agent any, src *messagingx.Ticker) {
	fmt.Printf("OnTick()  -> %v : ticker:%v\n", DispatchName(agent), DispatchName(src))
}

func (t *testDispatch) OnMessage(agent any, msg *messagingx.Message, src *messagingx.Channel) {
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
	case messagingx.Agent:
		return ptr.Uri()
	case *messagingx.Ticker:
		return ptr.Name()
	case *messagingx.Channel:
		return ptr.Name()
	case *messagingx.Message:
		return ptr.Event()
	case *aspect.Status:
		return ptr.String()
	default:
		return fmt.Sprintf("%v", reflect.TypeOf(t))
	}
}
