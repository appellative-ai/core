package httpx

import (
	"context"
	"fmt"
	"net/http"
)

func ExampleBuildChainT() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChainT(do1{}, do2{}, do3{}, do4{})
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

func ExampleBuildChain_AbbreviatedT() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChainT(do1{}, do2{}, do3Fail{}, do4{})
	ex(req)

	//Output:
	//test: Do1() -> request
	//test: Do2() -> request
	//test: Do3() -> request
	//test: Do3() -> response
	//test: Do2() -> response
	//test: Do1() -> response

}

type do1 struct{}

func (d do1) Link(next Exchange) Exchange {
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

type do2 struct{}

func (d do2) Link(next Exchange) Exchange {
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

type do3 struct{}

func (d do3) Link(next Exchange) Exchange {
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

type do3Fail struct{}

func (d do3Fail) Link(next Exchange) Exchange {
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

type do4 struct{}

func (d do4) Link(next Exchange) Exchange {
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
