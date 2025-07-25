package std

import (
	"errors"
	"fmt"
	"net/http"
)

func ExampleStatus() {
	s := NewStatus(http.StatusTeapot, "agent\test", errors.New("this is an error"))

	fmt.Printf("test: NewStatus() -> %v\n", s)

	fmt.Printf("test: NewStatus() -> %v\n", StatusOK)

	fmt.Printf("test: NewStatus() -> %v\n", StatusNotFound)

	//Output:
	//test: NewStatus() -> I'm A Teapot - this is an error
	//test: NewStatus() -> OK
	//test: NewStatus() -> Not Found

}
