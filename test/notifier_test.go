package test

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
)

func ExampleNewNotifier() {
	n := NewNotifier()

	n.Notify(aspect.StatusNotFound())
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Status())

	n.Reset()
	n.Notify(aspect.StatusNoContent())
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Status())

	//Output:
	//test: NewNotifier() -> [status:Not Found]
	//test: NewNotifier() -> [status:No Content]

}
