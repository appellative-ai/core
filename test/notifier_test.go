package test

import (
	"errors"
	"fmt"
)

func ExampleNewNotifier() {
	n := NewNotifier()

	n.Notify(errors.New("error: not found"))
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Error())

	n.Reset()
	n.Notify(errors.New("error: no content"))
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Error())

	//Output:
	//test: NewNotifier() -> [status:Not Found]
	//test: NewNotifier() -> [status:No Content]

}
