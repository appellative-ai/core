package httpx

import (
	"fmt"
	"net/http"
)

func ExampleConfigMessage() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := func(r *http.Request) (*http.Response, error) {
		return NewResponse(http.StatusTeapot, nil, nil), nil
	}
	m := NewConfigExchangeMessage(func(r *http.Request) (*http.Response, error) {
		return NewResponse(http.StatusTeapot, nil, nil), nil
	}, "")

	var name string
	ex, name, _ = ConfigExchangeContent(m)
	resp, err := ex(req)
	fmt.Printf("test: ConfigExchangeContent() -> [name:%v] [status:%v] [err:%v]\n", name, resp.StatusCode, err)

	m = NewConfigExchangeMessage(ex, "test-exchange")
	ex, name, _ = ConfigExchangeContent(m)
	resp, err = ex(req)
	fmt.Printf("test: ConfigExchangeContent() -> [name:%v] [status:%v] [err:%v]\n", name, resp.StatusCode, err)

	//Output:
	//test: ConfigExchangeContent() -> [name:default] [status:418] [err:<nil>]
	//test: ConfigExchangeContent() -> [name:test-exchange] [status:418] [err:<nil>]

}
