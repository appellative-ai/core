package messaging

import (
	"fmt"
	"net/http"
)

func ExampleNewMessage() {
	m := NewMessage("channel", StartupEvent)

	fmt.Printf("test: NewMessage() -> [%v]\n", m)

	//Output:
	//test: NewMessage() -> [[chan:channel] [from:] [to:] [common:core:event/startup]]

}

func ExampleMapMessage() {
	cfg := make(map[string]string)
	cfg["key1"] = "value1"
	cfg["key2"] = "value2"
	m := NewMapMessage(cfg)

	t, status := MapContent(m)
	fmt.Printf("test: MapContent() -> [%v] [status:%v]\n", t, status)

	m.SetContent(ContentTypeMap, "", "invalid content")
	t2, status2 := MapContent(m)
	fmt.Printf("test: MapContent() -> [%v] [status:%v]\n", t2, status2)

	//Output:
	//test: MapContent() -> [map[key1:value1 key2:value2]] [status:OK]
	//test: MapContent() -> [map[]] [status:Invalid Content [error: content value type: string is not of generic type: map[string]string]]

}

func ExampleStatusMessage() {
	m := NewStatusMessage(NewStatus(http.StatusTeapot, nil), ConfigEvent)

	status, event, result := StatusContent(m)
	fmt.Printf("test: StatusContent() -> [%v] [%v] [status:%v]\n", status, event, result)

	m.SetContent(ContentTypeStatus, "", "invalid content")
	status2, event2, result2 := StatusContent(m)
	fmt.Printf("test: StatusContent() -> [%v] [%v] [status:%v]\n", status2, event2, result2)

	//Output:
	//test: StatusContent() -> [I'm A Teapot] [common:core:event/config] [status:OK]
	//test: StatusContent() -> [<nil>] [] [status:Invalid Content [error: content value type: string is not of generic type: *messaging.Status]]

}
