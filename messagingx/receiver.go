package messagingx

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"net/http"
	"time"
)

type DoneFunc func(msg *Message) bool

func NewReceiverReplyTo(reply chan *Message) Handler {
	return func(msg *Message) {
		// needed when a timeout is reached, the channel is closed, and there is a pending send
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered in messaging.NewReceiverReplyTo() : %v\n", r)
			}
		}()
		if msg != nil {
			reply <- msg
		}
	}
}

// Receiver - receives reply messages and forwards to a function which will return true if the receiving is complete. The interval
// bounds the time spent receiving, and result status is sent on the status channel.
func Receiver(interval time.Duration, reply <-chan *Message, result chan<- *aspect.Status, done DoneFunc) {
	tick := time.Tick(interval)
	status := aspect.StatusOK()
	start := time.Now().UTC()

	if interval <= 0 || reply == nil || result == nil || done == nil {
		return
	}
	defer func() {
		result <- status
	}()
	for {
		select {
		case <-tick:
			status = aspect.NewStatusDuration(http.StatusGatewayTimeout, time.Since(start))
			return
		case msg, open := <-reply:
			if !open {
				status = aspect.NewStatusDuration(http.StatusInternalServerError, time.Since(start))
				return
			}
			if done(msg) {
				status = aspect.NewStatusDuration(http.StatusOK, time.Since(start))
				return
			}
		default:
		}
	}
}

func DrainAndClose(duration time.Duration, c chan *Message) {
	tick := time.Tick(time.Second * 10)
	for {
		select {
		case <-tick:
			close(c)
			return
		case <-c:
			close(c)
			return
		default:
		}
	}
}
