package host

import (
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"net/http"
	"net/http/httptest"
)

func Exchange(w http.ResponseWriter, r *http.Request, handler rest.Exchange) {
	httpx.AddRequestId(r)
	if handler == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, _ := handler(r)
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}

func ExampleHost() {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8081/github/advanced-go/search:google?q=golang", nil)

	Exchange(r, req, func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK}, nil
	})

	fmt.Printf("test: Exchange() -> [resp:%v]\n", r.Result().StatusCode)

	//Output:
	//test: Exchange() -> [resp:200]

}

func ExampleNewEndpoint() {

}

/*
func ExampleExchangeHandler() {
	e := NewEndpoint2(nil)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	e.Exchange(rec, req)

	fmt.Printf("test: ExchangeHandler() -> [%v]\n", req.URL.String())

	//Output:
	//test: ExchangeHandler() -> [https://www.google.com/search?q=golang]

}


*/
