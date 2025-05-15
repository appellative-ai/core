package eventing

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNewAgent() {
	a := newAgent()
	status := messaging.NewStatus(http.StatusTeapot, errors.New("test error message"))
	a.Notify(status)

	a.AddActivity(ActivityEvent{
		Agent:   nil,
		Event:   "eventing",
		Source:  "source",
		Content: nil,
	})

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [core:agent/eventing]

}
