package httpx

import (
	"context"
	"fmt"
	"net/http"
)

func ExampleLinkT() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := LinkT(Do1, Do2, Do3, Do4)
	ex(req)

	//Output:
	//test: Do1() -> request
	//test: Do2() -> request
	//test: Do3() -> request
	//test: Do4() -> request
	//test: Do4() -> response
	//test: Do3() -> response
	//test: Do2() -> response
	//test: Do1() -> response

}

func ExampleLinkT_Abbreviated() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := LinkT(Do1, Do2, Do3Fail, Do4)
	ex(req)

	//Output:
	//test: Do1() -> request
	//test: Do2() -> request
	//test: Do3() -> request
	//test: Do3() -> response
	//test: Do2() -> response
	//test: Do1() -> response

}

func Do1(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do1() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do1() -> response\n")
		return
	}
}

func Do2(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do2() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do2() -> response\n")
		return
	}
}

func Do3(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do3() -> request\n")
		//fmt.Printf("test: Do3() -> response\n")
		//return &http.Response{StatusCode: http.StatusBadRequest}, nil
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do3() -> response\n")
		return
	}
}

func Do3Fail(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do3() -> request\n")
		fmt.Printf("test: Do3() -> response\n")
		return &http.Response{StatusCode: http.StatusBadRequest}, nil
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do3() -> response\n")
		return
	}
}

func Do4(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do4() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do4() -> response\n")
		return
	}
}

/*
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


*/
