package messaging

import (
	"fmt"
	"github.com/appellative-ai/core/std"
	"time"
)

func ExampleNewMessage() {
	m := NewMessage("channel", StartupEvent)

	fmt.Printf("test: NewMessage() -> [%v]\n", m)

	//Output:
	//test: NewMessage() -> [[chan:channel] [from:] [to:[]] [common:core:event/startup]]

}

func ExampleSetReply() {
	a := newControlAgent("test:agent/example", nil)
	a.run()
	m := NewMessage(ChannelControl, "test:agent/test")

	m.SetReply(nil)
	m.Reply(NewStatusMessage(std.StatusOK, ""))

	m.SetReply(m)
	m.Reply(NewStatusMessage(std.StatusOK, ""))

	m.SetReply(func(m *Message) {
		fmt.Printf("test: SetReply() -> %v\n", m)
	})
	m.Reply(NewStatusMessage(std.StatusNotFound, ""))

	m.SetReply(a)
	m.Reply(NewStatusMessage(std.StatusOK, ""))

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

	m.DeleteCareOf()
	fmt.Printf("test: CareOf() -> [%v]\n", m.CareOf())

	//Output:
	//test: CareOf() -> [test:agent/one]
	//test: CareOf() -> []

}
