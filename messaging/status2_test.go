package messaging

import (
	"errors"
	"fmt"
	"net/http"
)

func ExampleStatus2() {
	s := NewStatus2(http.StatusTeapot, "agent\test", errors.New("this is an error"))

	fmt.Printf("test: NewStatus() -> %v\n", s)

	fmt.Printf("test: NewStatus() -> %v\n", Status2OK)

	fmt.Printf("test: NewStatus() -> %v\n", Status2NotFound)

	//Output:
	//test: NewStatus() -> I'm A Teapot - this is an error
	//test: NewStatus() -> OK
	//test: NewStatus() -> Not Found

}
