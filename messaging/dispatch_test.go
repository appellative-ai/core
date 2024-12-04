package messaging

import "fmt"

func ExampleTraceDispatch_Channel() {
	d := NewTraceDispatcher(nil, EmissaryChannel)
	channel := ""
	event := ""

	d.Trace(nil, channel, event, "")
	fmt.Printf("test: Trace() -> [channel:%v]\n", channel)

	channel = EmissaryChannel
	d.Trace(nil, channel, event, "")
	fmt.Printf("test: Trace() -> [channel:%v]\n", channel)

	//Output:
	//test: Trace() -> [channel:]
	//trace -> 2024-11-24T18:40:08.606Z [emissary] [] [<nil>]
	//test: Trace() -> [channel:emissary] []

}

func ExampleTraceDispatch_Event() {
	d := NewTraceDispatcher([]string{ShutdownEvent, StartupEvent}, "")
	channel := ""
	event := ""

	d.Trace(nil, channel, event, "")
	fmt.Printf("test: Trace() -> [%v]\n", event)

	event = ShutdownEvent
	d.Trace(nil, channel, event, "")
	fmt.Printf("test: Trace() -> [%v]\n", event)

	event = ObservationEvent
	d.Trace(nil, channel, event, "")
	fmt.Printf("test: Trace() -> [channel:%v] [%v]\n", channel, event)

	//Output:
	//test: Trace() -> []
	//trace -> 2024-11-24T18:46:04.697Z [] [event:shutdown] [<nil>]
	//test: Trace() -> [event:shutdown]
	//test: Trace() -> [channel:] [event:observation]

}
