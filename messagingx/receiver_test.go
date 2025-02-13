package messagingx

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	//"github.com/behavioral-ai/aspect/aspect"

	//"github.com/behavioral-ai/aspect/aspect"
	"time"
)

func ExampleReceiver_Timeout() {
	var result *aspect.Status
	duration := time.Second * 2
	status := make(chan *aspect.Status, 1)
	reply := make(chan *Message, 16)

	go Receiver(duration, reply, status, func(msg *Message) bool { return true })
	result = <-status
	fmt.Printf("test: Receiver() -> [status:%v] [duration > %v:%v]\n", result, duration, result.Duration > duration)

	close(status)
	close(reply)

	//Output:
	//test: Receiver() -> [status:Timeout] [duration > 2s:true]

}

func ExampleReceiver_OK() {
	var result *aspect.Status
	duration := time.Second * 2
	status := make(chan *aspect.Status)
	reply := make(chan *Message, 16)

	go Receiver(duration, reply, status, func(msg *Message) bool {
		fmt.Printf("test: Receiver() - in Done()\n")
		return true
	})
	reply <- NewMessage(channelNone, "to", "from", "event", nil)
	result = <-status
	fmt.Printf("test: Receiver() -> [status:%v] [duration:%v]\n", result, result.Duration)

	close(status)
	close(reply)

	//Output:
	//test: Receiver() - in Done()
	//test: Receiver() -> [status:OK] [duration:0s]

}

func ExampleReceiver_Closed() {
	var result *aspect.Status
	duration := time.Second * 5
	status := make(chan *aspect.Status, 1)
	reply := make(chan *Message, 16)

	go Receiver(duration, reply, status, func(msg *Message) bool {
		fmt.Printf("test: Receiver() - in Done()\n")
		return true
	})
	close(reply)
	result = <-status
	fmt.Printf("test: Receiver() -> [status:%v] [duration:%v]\n", result, result.Duration)

	close(status)

	//Output:
	//test: Receiver() -> [status:Internal Error] [duration:0s]

}
