package messaging

import (
	"fmt"
	"github.com/behavioral-ai/core/fmtx"
	"time"
)

func _ExampleTicker() {
	t := NewTicker("messagingtest-ticker", time.Second*2)
	ctrl := make(chan *Message)

	go tickerRun(ctrl, t)
	time.Sleep(time.Second * 20)

	ctrl <- ShutdownMessage
	time.Sleep(time.Second * 2)

	//Output:
	//messagingtest: Ticker() -> 2024-07-11T14:39:57.164Z
	//messagingtest: Ticker() -> 2024-07-11T14:39:59.164Z
	//messagingtest: Ticker() -> 2024-07-11T14:40:04.182Z
	//messagingtest: Ticker() -> 2024-07-11T14:40:09.180Z
	//messagingtest: Ticker() -> 2024-07-11T14:40:11.193Z
	//messagingtest: Ticker() -> 2024-07-11T14:40:13.184Z

}

func tickerRun(ctrl <-chan *Message, t *Ticker) {
	count := 0
	t.Start(0)
	for {
		select {
		case <-t.ticker.C:
			fmt.Printf("messagingtest: Ticker() -> %v\n", fmtx.FmtRFC3339Millis(time.Now().UTC()))
			count++
			if count == 2 {
				t.Start(time.Second * 5)
			}
			if count == 4 {
				t.Reset()
			}
		case msg := <-ctrl:
			switch msg.Name() {
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

	fmt.Printf("messagingtest: IsFinalized() -> [finalized:%v]\n", t.IsFinalized())
	fmt.Printf("messagingtest: Stopped() -> %v\n", t.IsStopped())

	//Output:
	//messagingtest: IsFinalized() -> [finalized:true]
	//messagingtest: Stopped() -> true
}

func ExampleTicker_IsFinalized_False() {
	t := NewPrimaryTicker(time.Second * 5)
	//fmt.Printf("messagingtest: Stopped() -> %v\n", t.IsStopped())

	go func() {
		time.Sleep(time.Second * 20)
		t.Stop()
	}()

	fmt.Printf("messagingtest: IsFinalized() -> [finalized:%v]\n", t.IsFinalized())
	fmt.Printf("messagingtest: Stopped() -> %v\n", t.IsStopped())

	//Output:
	//messagingtest: IsFinalized() -> [finalized:false]
	//messagingtest: Stopped() -> false
}


*/
