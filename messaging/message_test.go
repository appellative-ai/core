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

func ExampleConfigMessage() {
	cfg := make(map[string]string)
	cfg["key1"] = "value1"
	cfg["key2"] = "value2"
	m := NewConfigMapMessage(cfg)

	m2 := ConfigMapContent(NewMessage(Master, ShutdownEvent))
	fmt.Printf("test: NewConfigMessage() -> [%v]\n", m2)

	fmt.Printf("test: NewConfigMessage() -> [%v]\n", ConfigMapContent(m))

	//Output:
	//test: NewConfigMessage() -> [map[]]
	//test: NewConfigMessage() -> [map[key1:value1 key2:value2]]

}

func ExampleStatusMessage() {
	m := NewStatusMessage(NewStatus(http.StatusTeapot), ConfigEvent)

	status, event := StatusContent(m)
	fmt.Printf("test: NewStatusMessage() -> [%v] [%v]\n", status, event)

	//Output:
	//test: NewStatusMessage() -> [I'm A Teapot] [event:config]

}

func ExampleSubscriptionMessage() {
	m := NewSubscriptionCreateMessage("create-to", "create-from", SubscriptionCreateEvent)
	s, ok := SubscriptionCreateContent(m)
	fmt.Printf("test: NewSubscriptionCreateMessage() -> [%v] [%v] [%v] [%v]\n", m.To(), m.Event(), s, ok)

	m = NewSubscriptionCancelMessage("cancel-to", "cancel-from", SubscriptionCancelEvent)
	s, ok = SubscriptionCancelContent(m)
	fmt.Printf("test: NewSubscriptionCancelMessage() -> [%v] [%v] [%v] [%v]\n", m.To(), m.Event(), s, ok)

	//Output:
	//test: NewSubscriptionCreateMessage() -> [create-to] [event:subscription-create] [{event:subscription-create create-from}] [true]
	//test: NewSubscriptionCancelMessage() -> [cancel-to] [event:subscription-cancel] [{event:subscription-cancel cancel-from}] [true]

}
