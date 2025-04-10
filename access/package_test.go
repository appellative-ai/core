package access

import (
	"fmt"
	"net/http"
	"time"
)

func ExampleLogAccess() {
	var start time.Time //:= time.Now().UTC()
	r, _ := http.NewRequest(http.MethodGet, "/github/advanced-go/example-domain/activity:entry", nil)
	r.Host = "localhost:8080"

	s := DefaultFormat(nil, IngressTraffic, start, time.Millisecond*2000, "ingress-route", r, &http.Response{StatusCode: http.StatusGatewayTimeout, Status: "Gateway Timeout"}, Threshold{Timeout: time.Millisecond * 1500})
	fmt.Printf("test: log() -> %v\n", s)

	s = DefaultFormat(nil, IngressTraffic, start, time.Millisecond*750, "", r, &http.Response{StatusCode: http.StatusTooManyRequests, Status: "Too Many Requests"}, Threshold{RateLimit: 50})
	fmt.Printf("test: log() -> %v\n", s)

	s = DefaultFormat(nil, EgressTraffic, start, time.Millisecond*345, "egress-route", r, &http.Response{StatusCode: 200, Status: "OK"}, Threshold{Redirect: 10})
	fmt.Printf("test: log() -> %v\n", s)

	//Output:
	//test: log() -> {"traffic":"ingress", "start":0001-01-01T00:00:00.000Z, "duration":2000, "route":"ingress-route", "request-id":null, "protocol":"HTTP/1.1", "method":"GET", "host":"localhost:8080", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":504, "encoding":null, "bytes":0, "timeout":1500, "rate-limit":-1, "redirect":-1 }
	//test: log() -> {"traffic":"ingress", "start":0001-01-01T00:00:00.000Z, "duration":750, "route":null, "request-id":null, "protocol":"HTTP/1.1", "method":"GET", "host":"localhost:8080", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":429, "encoding":null, "bytes":0, "timeout":-1, "rate-limit":50, "redirect":5 }
	//test: log() -> {"traffic":"egress", "start":0001-01-01T00:00:00.000Z, "duration":345,  "route":"egress-route", "request-id":null, "protocol":"HTTP/1.1", "method":"GET", "host":"localhost:8080", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":200, "encoding":null, "bytes":0, "timeout":-1, "rate-limit":-1, "redirect":10 }

}
