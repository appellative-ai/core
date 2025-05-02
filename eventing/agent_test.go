package eventing

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNewAgent() {
	a := newAgent()
	status := messaging.NewStatusError(http.StatusTeapot, errors.New("test error message"), NamespaceName)
	a.Notify(status)

	a.AddActivity(ActivityEvent{
		Agent:   nil,
		Event:   "eventing",
		Source:  "source",
		Content: nil,
	})

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/eventing]

}
