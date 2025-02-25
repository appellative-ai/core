package test

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNewNotifier() {
	n := NewNotifier()

	n.Notify(messaging.NewStatusError(http.StatusNotFound, errors.New("error: not found"), "", nil))
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Error())

	n.Reset()
	n.Notify(messaging.NewStatusError(http.StatusNoContent, errors.New("error: no content"), "", nil))
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Error())

	//Output:
	//test: NewNotifier() -> [status:error: not found]
	//test: NewNotifier() -> [status:error: no content]

}
