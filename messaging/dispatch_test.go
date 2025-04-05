package messaging

import (
	"fmt"
)

func ExampleTraceDispatch_Channel() {
	d := NewFilteredTraceDispatcher(nil, Emissary)
	channel := ""
	event := ""

	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [channel:%v]\n", channel)

	channel = Emissary
	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [channel:%v]\n", channel)

	//Output:
	//test: Dispatch() -> [channel:]
	//trace -> 2024-11-24T18:40:08.606Z [<nil>] [emissary] []
	//test: Dispatch() -> [channel:emissary]

}

func ExampleTraceDispatch_Event() {
	d := NewFilteredTraceDispatcher([]string{ShutdownEvent, StartupEvent}, "")
	channel := ""
	event := ""

	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [%v]\n", event)

	event = ShutdownEvent
	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [%v]\n", event)

	event = ObservationEvent
	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [channel:%v] [%v]\n", channel, event)

	//Output:
	//test: Dispatch() -> []
	//trace -> 2024-11-24T18:46:04.697Z [<nil>] [] [eventing:shutdown]
	//test: Dispatch() -> [eventing:shutdown]
	//test: Dispatch() -> [channel:] [eventing:observation]

}
