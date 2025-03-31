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
	})

	ex, _ = ConfigExchangeContent(m)
	resp, err := ex(req)
	fmt.Printf("test: ConfigExchangeContent() -> [status:%v] [err:%v]\n", resp.StatusCode, err)

	m = NewConfigExchangeMessage(ex)
	ex, _ = ConfigExchangeContent(m)
	resp, err = ex(req)
	fmt.Printf("test: ConfigExchangeContent() -> [status:%v] [err:%v]\n", resp.StatusCode, err)

	//Output:
	//test: ConfigExchangeContent() -> [status:418] [err:<nil>]
	//test: ConfigExchangeContent() -> [status:418] [err:<nil>]

}
