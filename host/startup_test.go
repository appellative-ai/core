package host

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messagingx"
	msg2 "github.com/behavioral-ai/core/test"
	"net/http"
	"time"
)

func emptyRun(uri string, ctrl, data <-chan *messagingx.Message, state any) {
}
func testRegister(ex *messagingx.Exchange, uri string, cmd chan *messagingx.Message) error {
	a := msg2.NewAgent(uri)
	ex.Register(a) //.NewMailboxWithCtrl(uri, false, cmd, data))
	return nil
}

var start time.Time

func ExampleCreateToSend() {
	uriNone := "startup/none"
	uriOne := "startup/one"
	ex := messagingx.NewExchange()

	a := msg2.NewAgent(uriNone)
	err := ex.Register(a)
	if err != nil {
		fmt.Printf("test: NewAgent(%v) -> [err:%v]\n", uriNone, err)
	}

	a = msg2.NewAgent(uriOne)
	err = ex.Register(a)
	if err != nil {
		fmt.Printf("test: NewAgent(%v) -> [err:%v]\n", uriOne, err)
	}
	m := createToSend(ex, nil, nil)
	msg := m[uriNone]
	fmt.Printf("test: createToSend(nil,nil) -> [to:%v] [from:%v]\n", msg.To(), msg.From())

	//Output:
	//test: createToSend(nil,nil) -> [to:startup/none] [from:github/advanced-go/aspect/host:Startup]

}

func ExampleStartup_Success() {
	uri1 := "github/startup/good"
	uri2 := "github/startup/bad"
	uri3 := "github/startup/depends"

	ex := messagingx.NewExchange()
	start = time.Now()

	c := make(chan *messagingx.Message, 16)
	testRegister(ex, uri1, c)
	go startupGood(c)

	c = make(chan *messagingx.Message, 16)
	testRegister(ex, uri2, c)
	go startupBad(c)

	c = make(chan *messagingx.Message, 16)
	testRegister(ex, uri3, c)
	go startupDepends(c, nil)

	status := startup(ex, time.Second*2, nil)

	fmt.Printf("test: Startup() -> [%v]\n", status)

	//Output:
	//startup successful: [github/startup/bad] : 0s
	//startup successful: [github/startup/depends] : 0s
	//startup successful: [github/startup/good] : 0s
	//test: Startup() -> [true]

}

func ExampleStartup_Failure() {
	uri1 := "github/startup/good"
	uri2 := "github/startup/bad"
	uri3 := "github/startup/depends"
	ex := messagingx.NewExchange()

	start = time.Now()

	c := make(chan *messagingx.Message, 16)
	testRegister(ex, uri1, c)
	go startupGood(c)

	c = make(chan *messagingx.Message, 16)
	testRegister(ex, uri2, c)
	go startupBad(c)

	c = make(chan *messagingx.Message, 16)
	testRegister(ex, uri3, c)
	go startupDepends(c, errors.New("startup failure error message"))

	status := startup(ex, time.Second*2, nil)

	fmt.Printf("test: Startup() -> [%v]\n", status)

	//Output:
	//error: startup failure [startup failure error message]
	//test: Startup() -> [false]

}

func startupGood(c chan *messagingx.Message) {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			messagingx.SendReply(msg, aspect.NewStatusDuration(http.StatusOK, time.Since(start)))
		default:
		}
	}
}

func startupBad(c chan *messagingx.Message) {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			time.Sleep(time.Second + time.Millisecond*100)
			messagingx.SendReply(msg, aspect.NewStatusDuration(http.StatusOK, time.Since(start)))
		default:
		}
	}
}

func startupDepends(c chan *messagingx.Message, err error) {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			if err != nil {
				time.Sleep(time.Second)
				s := aspect.NewStatusDuration(0, time.Since(start))
				s.Err = err
				messagingx.SendReply(msg, s)
			} else {
				time.Sleep(time.Second + (time.Millisecond * 900))
				messagingx.SendReply(msg, aspect.NewStatusDuration(http.StatusOK, time.Since(start)))
			}

		default:
		}
	}
}
