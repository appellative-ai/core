package rest

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"net/http"
)

func ExampleUpdateExchange() {
	UpdateExchange("", nil, nil)
	fmt.Printf("test: UpdateExchange() -> nil Review\n")

	//var ex Exchange
	ex := func(r *http.Request) (*http.Response, error) {
		fmt.Printf("test: UpdateExchange() -> original\n")
		return nil, nil
	}
	UpdateExchange("", (*Exchange)(&ex), nil)
	fmt.Printf("test: UpdateExchange() -> nil message\n")

	m := messaging.NewMessage(messaging.ChannelControl, "test-message")
	UpdateExchange("", (*Exchange)(&ex), m)
	fmt.Printf("test: UpdateExchange() -> invalid content type\n")

	var ex2 Exchange
	ex2 = func(r *http.Request) (*http.Response, error) {
		fmt.Printf("test: UpdateExchange() -> updated\n")
		return nil, nil
	}
	m = NewExchangeMessage(ex2) //func(r *http.Request) (*http.Response, error) {
	//fmt.Printf("test: UpdateExchange() -> updated\n")
	//return nil, nil
	//})
	original := ex
	UpdateExchange("", (*Exchange)(&ex), m)
	original(nil)
	ex(nil)

	//Output:
	//test: UpdateExchange() -> nil Review
	//test: UpdateExchange() -> nil message
	//test: UpdateExchange() -> invalid content type
	//test: UpdateExchange() -> original
	//test: UpdateExchange() -> updated

}
