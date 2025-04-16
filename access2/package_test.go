package access2

import (
	"net/http"
	"time"
)

func ExampleLog_Default() {
	start := time.Now().UTC()

	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com/search?q=golang", nil)
	resp := &http.Response{StatusCode: http.StatusTeapot, ContentLength: 12345}
	resp.Header = make(http.Header)
	resp.Header.Add(ContentEncoding, "gzip")
	resp.Header.Add(XCached, "true")
	t := Threshold{Timeout: time.Millisecond * 456, RateLimit: 100, Redirect: 8}
	Log(nil, EgressTraffic, start, time.Millisecond*1500, "test-route", req, resp, t)

	//Output:
	//fail

}
