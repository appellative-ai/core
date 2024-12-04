package test

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
)

func ExampleNewNotifier() {
	n := NewNotifier()

	n.Notify(core.StatusNotFound())
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Status())

	n.Reset()
	n.Notify(core.StatusNoContent())
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Status())

	//Output:
	//test: NewNotifier() -> [status:Not Found]
	//test: NewNotifier() -> [status:No Content]

}
