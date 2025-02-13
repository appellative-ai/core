package messaging

import "fmt"

func ExampleNewMessage() {
	m := NewMessage("channel", "to", "from", StartupEvent)

	fmt.Printf("test: NewMessage() -> [%v]\n", m)

	//Output:
	//fail
}
