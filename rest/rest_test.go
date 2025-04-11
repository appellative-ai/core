package rest

import (
	"context"
	"fmt"
	"net/http"
)

func ExampleBuildChain_Link() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1Fn, do2Fn, do3Fn, do4Fn)
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

func ExampleBuildChain_Chainable() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, do2{}, do3{}, do4{})
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

func ExampleBuildChain_Any() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, do2Fn, do3{}, do4Fn)
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

func _ExampleBuildChain_Panic_Type() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, req, do3{}, do4Fn)
	ex(req)

	//Output:
	//fail
}

func _ExampleBuildChain_Panic_Nil() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, nil, do3{}, do4Fn)
	ex(req)

	//Output:
	//fail
}

func ExampleBuildChain_Abbreviated() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1Fn, do2Fn, do3FailFn, do4Fn)
	ex(req)

	//Output:
	//test: Do1() -> request
	//test: Do2() -> request
	//test: Do3() -> request
	//test: Do3() -> response
	//test: Do2() -> response
	//test: Do1() -> response

}

func do1Fn(next Exchange) Exchange {
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

func do2Fn(next Exchange) Exchange {
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

func do3Fn(next Exchange) Exchange {
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

func do3FailFn(next Exchange) Exchange {
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

func do4Fn(next Exchange) Exchange {
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

type do1 struct{}

func (d do1) Link(next Exchange) Exchange {
	return do1Fn(next)
}

type do2 struct{}

func (d do2) Link(next Exchange) Exchange {
	return do2Fn(next)
}

type do3 struct{}

func (d do3) Link(next Exchange) Exchange {
	return do3Fn(next)
}

type do3Fail struct{}

func (d do3Fail) Link(next Exchange) Exchange {
	return do3FailFn(next)
}

type do4 struct{}

func (d do4) Link(next Exchange) Exchange {
	return do4Fn(next)
}
