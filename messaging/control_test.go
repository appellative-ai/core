package messaging

import (
	"fmt"
	"time"
)

const (
	PingEvent        = "event:ping"
	ReconfigureEvent = "event:reconfigure"
)

func newAgentCtrlHandler(msg *Message) {
	fmt.Printf(fmt.Sprintf("test: NewControlAgent_CtrlHandler() -> %v\n", msg.Event()))
}

func ExampleNewControlAgent() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("test: NewControlAgent() -> [recovered:%v]\n", r)
		}
	}()
	//ex := NewExchange() //any(NewDirectory()).(*directory)
	//c := make(chan Message, 16)
	uri := "github.com/advanced-go/example-domain/activity"

	a, error0 := NewControlAgent(uri, newAgentCtrlHandler)
	if error0 != nil {
		fmt.Printf("test: NewControlAgent() -> [err:%v]\n", error0)
	}
	//status = a.Register(agentDir)
	//if !status.OK() {
	//	fmt.Printf("test: Register() -> [status:%v]\n", status)
	//}
	// 1 -10 Nanoseconds works for a direct send to a channel, sending via an controller2 needs a longer sleep time
	//d := time.Nanosecond * 10
	// Needed time.Nanoseconds * 50 for directory send with mutex
	// Needed time.Nanoseconds * 1 for directory send via sync.Map
	d := time.Nanosecond * 1
	a.Run()
	a.Message(NewMessage(ControlChannel, StartupEvent))
	//c <- Message{To: "", From: "", Event: aspect.StartupEvent, RelatesTo: "", Status: nil, Content: nil, ReplyTo: nil}
	time.Sleep(d)
	a.Message(NewMessage(ControlChannel, PauseEvent))
	//c <- Message{To: "", From: "", Event: aspect.PauseEvent, RelatesTo: "", Status: nil, Content: nil, ReplyTo: nil}
	time.Sleep(d)
	a.Message(NewMessage(ControlChannel, ResumeEvent))
	//c <- Message{To: "", From: "", Event: aspect.ResumeEvent, RelatesTo: "", Status: nil, Content: nil, ReplyTo: nil}
	time.Sleep(d)
	a.Message(NewMessage(ControlChannel, PingEvent))
	//c <- Message{To: "", From: "", Event: aspect.PingEvent, RelatesTo: "", Status: nil, Content: nil, ReplyTo: nil}
	time.Sleep(d)
	a.Message(NewMessage(ControlChannel, ReconfigureEvent))
	//c <- Message{To: "", From: "", Event: aspect.ReconfigureEvent, RelatesTo: "", Status: nil, Content: nil, ReplyTo: nil}
	time.Sleep(d)
	a.Shutdown() //.SendCtrl(Message{To: uri, From: "", Event: aspect.ShutdownEvent})
	//c <- Message{To: "", From: "", Event: aspect.ShutdownEvent, RelatesTo: "", Status: nil, Content: nil, ReplyTo: nil}
	time.Sleep(time.Millisecond * 100)

	a.Shutdown()
	// will panic
	//c <- Message{}

	//Output:
	//test: NewControlAgent_CtrlHandler() -> messaging:event.startup
	//test: NewControlAgent_CtrlHandler() -> messaging:event.pause
	//test: NewControlAgent_CtrlHandler() -> messaging:event.resume
	//test: NewControlAgent_CtrlHandler() -> event:ping
	//test: NewControlAgent_CtrlHandler() -> event:reconfigure
	//test: NewControlAgent_CtrlHandler() -> messaging:event.shutdown

}
