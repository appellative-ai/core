package eventing

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNotify() {
	s := messaging.NewStatus(http.StatusGatewayTimeout, errors.New("rate limiting"))
	s.WithLocation("resiliency:agent/operative")
	s.WithRequestId("123-request-id")

	OutputNotify(s)

	//Output:
	//notify-> 2025-02-26T15:34:45.784Z [resiliency:agent/operative] [core:messaging.status] [123-request-id] [Timeout] [rate limiting]

}

func ExampleNewStatus_Error() {
	s := messaging.NewStatus(http.StatusGatewayTimeout, errors.New("rate limited")) //"resiliency:agent/operative/agent1#us-west")
	s.WithLocation("test:agent")
	fmt.Printf("test: NewStatus_Error() -> [%v]\n", s)

	if _, ok := any(s).(NotifyEvent); ok {
		fmt.Printf("test: Event() -> [%v]\n", ok)

	}

	//Output:
	//test: NewStatus_Error() -> [Timeout [err:rate limited] [location:test:agent]]
	//test: Event() -> [true]

}

func ExampleNewStatus_Message() {
	s := messaging.NewStatus(http.StatusOK, "successfully change ticker duration")
	s.WithLocation("test:agent")
	fmt.Printf("test: NewStatus_Message() -> [%v]\n", s)

	if _, ok := any(s).(NotifyEvent); ok {
		fmt.Printf("test: Event() -> [%v]\n", ok)

	}

	//Output:
	//test: NewStatus_Message() -> [OK [msg:successfully change ticker duration] [location:test:agent]]
	//test: Event() -> [true]

}
