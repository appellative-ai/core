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

	m2 := ConfigMapContent(NewMessage(ChannelMaster, ShutdownEvent))
	fmt.Printf("test: NewConfigMessage() -> [%v]\n", m2)

	fmt.Printf("test: NewConfigMessage() -> [%v]\n", ConfigMapContent(m))

	//Output:
	//test: NewConfigMessage() -> [map[]]
	//test: NewConfigMessage() -> [map[key1:value1 key2:value2]]

}

func ExampleStatusMessage() {
	m := NewStatusMessage(NewStatus(http.StatusTeapot, nil), ConfigEvent)

	status, event := StatusContent(m)
	fmt.Printf("test: NewStatusMessage() -> [%v] [%v]\n", status, event)

	//Output:
	//test: NewStatusMessage() -> [I'm A Teapot] [event:config]

}
