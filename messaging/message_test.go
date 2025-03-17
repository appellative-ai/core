package messaging

import (
	"fmt"
	"net/http"
)

func ExampleNewMessage() {
	m := NewMessage("channel", StartupEvent)

	fmt.Printf("test: NewMessage() -> [%v]\n", m)

	//Output:
	//test: NewMessage() -> [[chan:channel] [from:] [to:] [event:startup]]

}

func ExampleNotifyMessage() {
	status := NewStatusMessage(http.StatusTeapot, "test message", "agent/test")
	m := NewNotifyMessage(status)
	e := EventContent(m)
	fmt.Printf("test: EventContent() -> [%v]\n", e)

	//Output:
	//test: EventContent() -> [I'm A Teapot [msg:test message] [agent:agent/test]]

}
