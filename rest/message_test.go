package rest

import (
	"fmt"
	"net/http"
)

func ExampleNewRouteMessage() {
	m := NewRouteMessage("test-route", "https://www.google.com/search?q=golang", func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	})
	r, ok := RouteContent(m)
	fmt.Printf("test: NewRouteMessage() [name:%v] [uri:%v] [ok:%v]\n", r.Name, r.Uri, ok)

	//Output:
	//test: NewRouteMessage() [name:test-route] [uri:https://www.google.com/search?q=golang] [ok:true]

}

func ExampleNewExchangeMessage() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	}
	m := NewExchangeMessage(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	})

	ex, _ = ExchangeContent(m)
	resp, err := ex(req)
	fmt.Printf("test: ExchangeContent() -> [status:%v] [err:%v]\n", resp.StatusCode, err)

	m = NewExchangeMessage(ex)
	ex, _ = ExchangeContent(m)
	resp, err = ex(req)
	fmt.Printf("test: ExchangeContent() -> [status:%v] [err:%v]\n", resp.StatusCode, err)

	//Output:
	//test: ExchangeContent() -> [status:418] [err:<nil>]
	//test: ExchangeContent() -> [status:418] [err:<nil>]

}
