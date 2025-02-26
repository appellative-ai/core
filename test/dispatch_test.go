package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleDispatchName() {
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(nil))

	a := NewAgent("agent-test")
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(a))

	t := messaging.NewTicker("ticker-test", 100)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(t))

	c := messaging.NewChannel("channel-test")
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(c))

	m := messaging.NewMessage(messaging.ControlChannel, "event-test")
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(m))

	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(messaging.StatusNotFound()))

	r := new(http.Response)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(r))

	//Output:
	//test: DispatchName() -> <nil>
	//test: DispatchName() -> agent-test
	//test: DispatchName() -> ticker-test
	//test: DispatchName() -> channel-test
	//test: DispatchName() -> event-test
	//test: DispatchName() -> *messaging.Status
	//test: DispatchName() -> *http.Response

}
