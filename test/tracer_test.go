package test

import "github.com/behavioral-ai/core/messagingx"

func ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test")
	DefaultTracer.Trace(nil, messagingx.MasterChannel, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	DefaultTracer.Trace(a, messagingx.EmissaryChannel, "event:shutdown", "")
	//fmt.Printf("\n")

	//Output:
	//test: Trace() -> <nil> : [master] [event:shutdown] [agent shutdown]
	//test: Trace() -> agent:test : [emissary] [event:shutdown] [agent shutdown]

}
