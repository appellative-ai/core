package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleNewAgent() {
	a := NewAgent("urn:any")
	if _, ok := any(a).(messaging.Agent); ok {
		fmt.Printf("test: opsAgent() -> ok\n")
	} else {
		fmt.Printf("test: opsAgent() -> fail\n")
	}

	//Output:
	//test: opsAgent() -> ok

}

func ExampleOld() {
	fmt.Printf("test")

	//Output:
	//test

}
