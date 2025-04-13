package host

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleHost() {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8081/github/advanced-go/search:google?q=golang", nil)

	Exchange3(r, req, func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK}, nil
	})

	fmt.Printf("test: Exchange() -> [resp:%v]\n", r.Result().StatusCode)

	//Output:
	//test: Exchange() -> [resp:200]

}
