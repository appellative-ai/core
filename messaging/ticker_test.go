package messaging

import (
	"fmt"
	"time"
)

func _ExampleTicker() {
	t := NewTicker("test-ticker", time.Second*2)
	ctrl := make(chan *Message)

	go tickerRun(ctrl, t)
	time.Sleep(time.Second * 20)

	ctrl <- Shutdown
	time.Sleep(time.Second * 2)

	//Output:
	//test: Ticker() -> 2024-07-11T14:39:57.164Z
	//test: Ticker() -> 2024-07-11T14:39:59.164Z
	//test: Ticker() -> 2024-07-11T14:40:04.182Z
	//test: Ticker() -> 2024-07-11T14:40:09.180Z
	//test: Ticker() -> 2024-07-11T14:40:11.193Z
	//test: Ticker() -> 2024-07-11T14:40:13.184Z

}

func tickerRun(ctrl <-chan *Message, t *Ticker) {
	count := 0
	t.Start(0)
	for {
		select {
		case <-t.ticker.C:
			fmt.Printf("test: Ticker() -> %v\n", FmtRFC3339Millis(time.Now().UTC()))
			count++
			if count == 2 {
				t.Start(time.Second * 5)
			}
			if count == 4 {
				t.Reset()
			}
		case msg := <-ctrl:
			switch msg.Event() {
			case ShutdownEvent:
				return
			default:
			}
		default:
		}
	}
}

/*
func ExampleTicker_IsFinalized_True() {
	t := NewPrimaryTicker(time.Second * 5)
	go func() {
		time.Sleep(time.Second * 6)
		t.Stop()
	}()

	fmt.Printf("test: IsFinalized() -> [finalized:%v]\n", t.IsFinalized())
	fmt.Printf("test: Stopped() -> %v\n", t.IsStopped())

	//Output:
	//test: IsFinalized() -> [finalized:true]
	//test: Stopped() -> true
}

func ExampleTicker_IsFinalized_False() {
	t := NewPrimaryTicker(time.Second * 5)
	//fmt.Printf("test: Stopped() -> %v\n", t.IsStopped())

	go func() {
		time.Sleep(time.Second * 20)
		t.Stop()
	}()

	fmt.Printf("test: IsFinalized() -> [finalized:%v]\n", t.IsFinalized())
	fmt.Printf("test: Stopped() -> %v\n", t.IsStopped())

	//Output:
	//test: IsFinalized() -> [finalized:false]
	//test: Stopped() -> false
}


*/
