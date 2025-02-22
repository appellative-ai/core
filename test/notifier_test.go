package test

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNewNotifier() {
	n := NewNotifier()

	n.Notify(messaging.NewStatusError(http.StatusNotFound, errors.New("error: not found")))
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Error())

	n.Reset()
	n.Notify(messaging.NewStatusError(http.StatusNoContent, errors.New("error: no content")))
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Error())

	//Output:
	//test: NewNotifier() -> [status:Not Found]
	//test: NewNotifier() -> [status:No Content]

}

func ExampleNotify() {
	Notify(messaging.NewStatusError(http.StatusNotFound, errors.New("error: not found")))

	//Output:
	//Not Found [error: not found]

}
