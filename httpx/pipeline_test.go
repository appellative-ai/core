package httpx

import (
	"context"
	"fmt"
	"net/http"
)

func do1(req *http.Request, next *Frame) (*http.Response, error) {
	fmt.Printf("test: do1() -> request\n")
	if next != nil {
		next.Fn(req, next.Next)

	}
	fmt.Printf("test: do1() -> response\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func do2(req *http.Request, next *Frame) (*http.Response, error) {
	fmt.Printf("test: do2() -> request\n")
	if next != nil {
		next.Fn(req, next.Next)
	}
	fmt.Printf("test: do2() -> response\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func do3(req *http.Request, next *Frame) (*http.Response, error) {
	fmt.Printf("test: do3() -> request\n")
	//return &httpx.Response{StatusCode: httpx.StatusOK}, nil
	if next != nil {
		next.Fn(req, next.Next)
	}
	fmt.Printf("test: do3() -> response\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func do4(req *http.Request, next *Frame) (*http.Response, error) {
	fmt.Printf("test: do4() -> request\n")
	if next != nil {
		next.Fn(req, next.Next)
	}
	fmt.Printf("test: do4() -> response\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func ExampleExchangePipeline_New() {
	p := NewExchangePipeline(do1, do2, do3, do4)
	for f := p.head; f != nil; f = f.Next {
		fmt.Printf("test: Head() -> [curr:%v] [next:%v]\n", f, f.Next)
	}

	//Output:
	//test: Head() -> [curr:do1] [next:do2]
	//test: Head() -> [curr:do2] [next:do3]
	//test: Head() -> [curr:do3] [next:do4]
	//test: Head() -> [curr:do4] [next:<nil>]

}

func ExampleExchangePipeline_Run() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	p := NewExchangePipeline(do1, do2, do3, do4)

	p.Run(req)

	//Output:
	//test: do1() -> request
	//test: do2() -> request
	//test: do3() -> request
	//test: do4() -> request
	//test: do4() -> response
	//test: do3() -> response
	//test: do2() -> response
	//test: do1() -> response

}
