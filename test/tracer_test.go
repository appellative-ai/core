package test

import "github.com/behavioral-ai/core/messaging"

func ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test")
	DefaultTracer.Trace(nil, messaging.MasterChannel, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	DefaultTracer.Trace(a, messaging.EmissaryChannel, "event:shutdown", "")
	//fmt.Printf("\n")

	//Output:
	//test: Trace() -> <nil> : [master] [event:shutdown] [agent shutdown]
	//test: Trace() -> agent:test : [emissary] [event:shutdown] [agent shutdown]

}
