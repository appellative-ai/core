package host

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleHost() {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8081/github/advanced-go/search:google?q=golang", nil)

	hostExchange(r, req, 0, func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK}, nil
	})

	fmt.Printf("test: hostExchange() -> [resp:%v]\n", r.Result().StatusCode)

	//Output:
	//test: hostExchange() -> [resp:200]

}
