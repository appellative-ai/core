package rest

import (
	"fmt"
	"net/http"
)

func ExampleNewRouteMessage() {
	m := NewRouteMessage("test-route", "https://www.google.com/search?q=golang", func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	})
	r, status := RouteContent(m)
	fmt.Printf("test: RouteContent() [name:%v] [uri:%v] [status:%v]\n", r.Name, r.Uri, status)

	//Output:
	//test: RouteContent() [name:test-route] [uri:https://www.google.com/search?q=golang] [status:OK]

}

func ExampleNewExchangeMessage() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	m := NewExchangeMessage(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	})

	ex, status := ExchangeContent(m)
	resp, err := ex(req)
	fmt.Printf("test: ExchangeContent() -> [status:%v] [code:%v] [err:%v]\n", status, resp.StatusCode, err)

	/*
		m = NewExchangeMessage(ex)
		ex, status = ExchangeContent(m)
		resp, err = ex(req)
		fmt.Printf("test: ExchangeContent() -> [status:%v] [err:%v]\n", resp.StatusCode, err)


	*/

	//Output:
	//test: ExchangeContent() -> [status:OK] [code:418] [err:<nil>]

}
