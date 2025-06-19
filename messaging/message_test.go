package messaging

import (
	"fmt"
	"net/http"
	"time"
)

func ExampleNewMessage() {
	m := NewMessage("channel", StartupEvent)

	fmt.Printf("test: NewMessage() -> [%v]\n", m)

	//Output:
	//test: NewMessage() -> [[chan:channel] [from:] [to:[]] [common:core:event/startup]]

}

func ExampleMapMessage() {
	cfg := make(map[string]string)
	cfg["key1"] = "value1"
	cfg["key2"] = "value2"
	m := NewMapMessage(cfg)

	t, status := MapContent(m)
	fmt.Printf("test: MapContent() -> [%v] [status:%v]\n", t, status)

	m.SetContent(ContentTypeMap, "invalid content")
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

	m.SetContent(ContentTypeStatus, "invalid content")
	status2, event2, result2 := StatusContent(m)
	fmt.Printf("test: StatusContent() -> [%v] [%v] [status:%v]\n", status2, event2, result2)

	//Output:
	//test: StatusContent() -> [I'm A Teapot] [common:core:event/config] [status:OK]
	//test: StatusContent() -> [<nil>] [] [status:Invalid Content [error: content value type: string is not of generic type: *messaging.Status]]

}

func ExampleAgentMessage() {
	a := newControlAgent("test:agent/example", nil)
	m := NewAgentMessage(a)

	a1, status := AgentContent(m)
	fmt.Printf("test: AgentContent() -> [%v] [status:%v]\n", a1, status)

	//Output:
	//test: AgentContent() -> [test:agent/example] [status:OK]

}

func ExampleSetReply() {
	a := newControlAgent("test:agent/example", nil)
	a.run()
	m := NewMessage(ChannelControl, "test:agent/test")

	m.SetReply(nil)
	m.Reply(NewStatusMessage(StatusOK(), ""))

	m.SetReply(m)
	m.Reply(NewStatusMessage(StatusOK(), ""))

	m.SetReply(func(m *Message) {
		fmt.Printf("test: SetReply() -> %v\n", m)
	})
	m.Reply(NewStatusMessage(StatusNotFound(), ""))

	m.SetReply(a)
	m.Reply(NewStatusMessage(StatusOK(), ""))

	time.Sleep(time.Second * 5)
	a.Message(ShutdownMessage)
	time.Sleep(time.Second * 5)

	//Output:
	//error: generic type is nil on call to messaging.SetReply
	//error: generic type: *messaging.Message, is invalid for messaging.SetReply
	//test: SetReply() -> [chan:ctrl] [from:] [to:[]] [common:core:event/status]
	//test: controlAgent.run() -> [chan:ctrl] [from:] [to:[]] [common:core:event/status]
	//test: controlAgent.run() -> [chan:ctrl] [from:] [to:[]] [common:core:event/shutdown]

}

func ExampleMessage_IsRecipient() {
	m := NewMessage(ChannelControl, ConfigEvent)
	m.AddTo("test:agent/one", "test:agent/two", "test:agent/three")

	name1 := ""
	ok := m.IsRecipient(name1)
	fmt.Printf("test: IsRecipient(\"%v\") -> [ok:%v]\n", name1, ok)

	name1 = "invalid"
	ok = m.IsRecipient(name1)
	fmt.Printf("test: IsRecipient(\"%v\") -> [ok:%v]\n", name1, ok)

	name1 = "test:agent/two"
	ok = m.IsRecipient(name1)
	fmt.Printf("test: IsRecipient(\"%v\") -> [ok:%v]\n", name1, ok)

	//Output:
	//test: IsRecipient("") -> [ok:false]
	//test: IsRecipient("invalid") -> [ok:false]
	//test: IsRecipient("test:agent/two") -> [ok:true]

}
func ExampleMessage_CareOf() {
	m := NewMessage(ChannelControl, ConfigEvent).SetCareOf("test:agent/one")
	m.AddTo("test:agent/two", "test:agent/three")

	fmt.Printf("test: CareOf() -> [%v]\n", m.CareOf())

	//Output:
	//test: CareOf() -> [test:agent/one]

}
