package test

import (
	"fmt"
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
