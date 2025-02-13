package test

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messagingx"
	"net/http"
)

func ExampleDispatchName() {
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(nil))

	a := NewAgent("agent-test")
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(a))

	t := messagingx.NewTicker("ticker-test", 100)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(t))

	c := messagingx.NewChannel("channel-test", false)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(c))

	m := messagingx.NewControlMessage("", "", "event-test")
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(m))

	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(aspect.StatusNotFound()))

	r := new(http.Response)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(r))

	//Output:
	//test: DispatchName() -> <nil>
	//test: DispatchName() -> agent-test
	//test: DispatchName() -> ticker-test
	//test: DispatchName() -> channel-test
	//test: DispatchName() -> event-test
	//test: DispatchName() -> Not Found
	//test: DispatchName() -> *http.Response

}
