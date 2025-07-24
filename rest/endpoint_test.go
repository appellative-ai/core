package rest

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"net/http"
)

func ExampleNewEndpoint() {
	e := NewEndpoint("/resource/test", nil, nil, nil)

	_, ok := any(e).(http.Handler)
	fmt.Printf("test: NewEndpoint() -> [%v] [%v] [ServeHTTP:%v]\n", e, e.Pattern(), ok)

	//Output:
	//test: NewEndpoint() -> [&{/resource/test <nil> <nil> <nil>}] [/resource/test] [ServeHTTP:true]

}

func ExampleNewExchangeMessage() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	m := messaging.NewConfigMessage(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	})

	ex, status := messaging.ConfigContent[func(r *http.Request) (*http.Response, error)](m) //ExchangeContent(m)
	resp, err := ex(req)
	fmt.Printf("test: ExchangeContent() -> [status:%v] [code:%v] [err:%v]\n", status, resp.StatusCode, err)

	//Output:
	//test: ExchangeContent() -> [status:true] [code:418] [err:<nil>]

}
