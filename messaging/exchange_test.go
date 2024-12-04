package messaging

import (
	"fmt"
	"time"
)

func emptyRun(uri string, ctrl, data <-chan *Message, state any) {
}

func emptyHandler(_ *Message) {}

func ExampleRegister() {
	testDir := NewExchange()

	fmt.Printf("test: Count() -> : %v\n", testDir.Count())

	uri1 := "urn:test:one"
	a := testDir.Get(uri1)
	fmt.Printf("test: Get(%v) -> : [agent:%v]\n", uri1, a)

	a1, _ := NewControlAgent(uri1, emptyHandler)
	err := testDir.Register(a1)
	fmt.Printf("test: Register(%v) -> : [err:%v]\n", uri1, err)

	fmt.Printf("test: Count() -> : %v\n", testDir.Count())
	m1 := testDir.Get(uri1)
	fmt.Printf("test: Get(%v) -> : [agent:%v]\n", uri1, m1.Uri())

	uri2 := "urn:test:two"
	a2, _ := NewControlAgent(uri2, emptyHandler)
	err = testDir.Register(a2)
	fmt.Printf("test: Register(%v) -> : [err:%v]\n", uri2, err)
	fmt.Printf("test: Count() -> : %v\n", testDir.Count())
	m2 := testDir.Get(uri2)
	fmt.Printf("test: Get(%v) -> : [agent:%v]\n", uri2, m2.Uri())

	fmt.Printf("test: List() -> : %v\n", testDir.List())

	//Output:
	//test: Count() -> : 0
	//test: Get(urn:test:one) -> : [agent:<nil>]
	//test: Register(urn:test:one) -> : [err:<nil>]
	//test: Count() -> : 1
	//test: Get(urn:test:one) -> : [agent:urn:test:one]
	//test: Register(urn:test:two) -> : [err:<nil>]
	//test: Count() -> : 2
	//test: Get(urn:test:two) -> : [agent:urn:test:two]
	//test: List() -> : [urn:test:one urn:test:two]

}

func ExampleRegisterError() {
	uri := "urn:agent007"
	ex := NewExchange()

	a, _ := NewControlAgent(uri, emptyHandler)
	err := ex.Register(a)
	fmt.Printf("test: Register(%v) -> [%v]\n", uri, err)

	err = ex.Register(a)
	fmt.Printf("test: Register(%v) -> [%v]\n", uri, err)

	//Output:
	//test: Register(urn:agent007) -> [<nil>]
	//test: Register(urn:agent007) -> [error: exchange.Register() agent already exists: [urn:agent007]]

}

func ExampleSendError() {
	uri := "urn:test"
	ex := NewExchange()

	fmt.Printf("test: Send(%v) -> : %v\n", uri, ex.Send(nil))
	fmt.Printf("test: Send(%v) -> : %v\n", uri, ex.Send(NewControlMessage("", "", "")))
	fmt.Printf("test: Send(%v) -> : %v\n", uri, ex.Send(NewControlMessage(uri, "", "")))

	//Output:
	//test: Send(urn:test) -> : error: controller2.Send() failed as message is nil
	//test: Send(urn:test) -> : error: controller2.Send() failed as the message To is empty or invalid : []
	//test: Send(urn:test) -> : error: controller2.Send() failed as the message To is empty or invalid : [urn:test]

}

func ExampleSend() {
	uri1 := "urn:agent-1"
	uri2 := "urn:agent-2"
	uri3 := "urn:agent-3"
	c := NewChannel("test", true) //make(chan *Message, 16)
	ex := NewExchange()

	a1 := newTestAgent(uri1, c, nil)
	ex.Register(a1)
	a2 := newTestAgent(uri2, c, nil)
	ex.Register(a2)
	a3 := newTestAgent(uri3, c, nil)
	ex.Register(a3)

	ex.Send(NewControlMessage(uri1, PkgPath, StartupEvent))
	ex.Send(NewControlMessage(uri2, PkgPath, StartupEvent))
	ex.Send(NewControlMessage(uri3, PkgPath, StartupEvent))

	time.Sleep(time.Second * 1)
	resp1 := <-c.C
	resp2 := <-c.C
	resp3 := <-c.C
	fmt.Printf("test: <- c -> : [%v] [%v] [%v]\n", resp1.To(), resp2.To(), resp3.To())
	c.Close()

	//Output:
	//test: <- c -> : [urn:agent-1] [urn:agent-2] [urn:agent-3]

}

func ExampleListCount() {
	uri1 := "urn:agent-1"
	uri2 := "urn:agent-2"
	ex := NewExchange()

	a1 := newTestAgent(uri1, nil, nil)
	ex.Register(a1)
	a2 := newTestAgent(uri2, nil, nil)
	ex.Register(a2)

	fmt.Printf("test: Count() -> : %v\n", ex.Count())
	fmt.Printf("test: List() -> : %v\n", ex.List())

	//Output:
	//test: Count() -> : 2
	//test: List() -> : [urn:agent-1 urn:agent-2]

}

func _ExampleExchangeOnShutdown() {
	uri1 := "urn:agent-1"
	uri2 := "urn:agent-2"
	ex := NewExchange()

	a1 := newTestAgent(uri1, nil, nil)
	ex.Register(a1)
	a2 := newTestAgent(uri2, nil, nil)
	ex.Register(a2)

	fmt.Printf("test: Get(%v) -> : %v\n", uri1, ex.Get(uri1))
	fmt.Printf("test: Get(%v) -> : %v\n", uri2, ex.Get(uri2))

	a1.running = true
	a1.Shutdown()

	a2.running = true
	a2.Shutdown()

	fmt.Printf("test: Get-Shutdown(%v) -> : %v\n", uri1, ex.Get(uri1))
	fmt.Printf("test: Get-Shutdown(%v) -> : %v\n", uri2, ex.Get(uri2))

	ex2 := NewExchange()
	ex.Register(a1)
	ex2.Register(a1)

	fmt.Printf("test: Get-Ex1(%v) -> : %v\n", uri1, ex.Get(uri1))
	fmt.Printf("test: Get-Ex2(%v) -> : %v\n", uri1, ex.Get(uri1))

	a1.running = true
	a1.Shutdown()
	fmt.Printf("test: Get-Ex1-Shutdown(%v) -> : %v\n", uri1, ex.Get(uri1))
	fmt.Printf("test: Get-Ex2-Shutdown(%v) -> : %v\n", uri1, ex.Get(uri1))

	//Output:
	//test: Get(urn:agent-1) -> : urn:agent-1
	//test: Get(urn:agent-2) -> : urn:agent-2
	//test: Get-Shutdown(urn:agent-1) -> : <nil>
	//test: Get-Shutdown(urn:agent-2) -> : <nil>
	//test: Get-Ex1(urn:agent-1) -> : urn:agent-1
	//test: Get-Ex2(urn:agent-1) -> : urn:agent-1
	//test: Get-Ex1-Shutdown(urn:agent-1) -> : <nil>
	//test: Get-Ex2-Shutdown(urn:agent-1) -> : <nil>

}
