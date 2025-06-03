package messaging

import "fmt"

func ExampleNewAgent() {
	a := NewExchangeAgent("/endpoint/resiliency")

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/exchange/endpoint/resiliency]

}
