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

	s := DefaultFormat(Origin{Region: "us", Zone: "zone", Host: r.Host, InstanceId: "id-012"}, IngressTraffic, start, time.Millisecond*2000, r, &http.Response{StatusCode: http.StatusGatewayTimeout, Status: "Gateway Timeout"}, Controller{Timeout: time.Millisecond * 1500, Code: ControllerTimeout})
	fmt.Printf("test: log() -> %v\n", s)

	s = DefaultFormat(Origin{Region: "us", Zone: "zone", Host: r.Host, InstanceId: "id-012"}, IngressTraffic, start, time.Millisecond*750, r, &http.Response{StatusCode: http.StatusTooManyRequests, Status: "Too Many Requests"}, Controller{RateLimit: "50", RateBurst: "5", Code: ControllerRateLimit})
	fmt.Printf("test: log() -> %v\n", s)

	s = DefaultFormat(Origin{Region: "us", Zone: "zone", Host: r.Host, InstanceId: "id-012"}, EgressTraffic, start, time.Millisecond*345, r, &http.Response{StatusCode: 200, Status: "OK"}, Controller{Percentage: "10", Code: ControllerRedirect})
	fmt.Printf("test: log() -> %v\n", s)

	//Output:
	//test: log() -> {"region":"us", "zone":"zone", "sub-zone":null, "instance-id":"id-012", "traffic":"ingress", "start":0001-01-01T00:00:00.000Z, "duration":2000, "request-id":null, "protocol":"HTTP/1.1", "method":"GET", "host":"localhost:8080", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":504, "encoding":null, "bytes":0, "timeout":1500, "rate-limit":-1, "rate-burst":-1, "cc":"TO" }
	//test: log() -> {"region":"us", "zone":"zone", "sub-zone":null, "instance-id":"id-012", "traffic":"ingress", "start":0001-01-01T00:00:00.000Z, "duration":750, "request-id":null, "protocol":"HTTP/1.1", "method":"GET", "host":"localhost:8080", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":429, "encoding":null, "bytes":0, "timeout":-1, "rate-limit":50, "rate-burst":5, "cc":"RL" }
	//test: log() -> {"region":"us", "zone":"zone", "sub-zone":null, "instance-id":"id-012", "traffic":"egress", "start":0001-01-01T00:00:00.000Z, "duration":345, "request-id":null, "protocol":"HTTP/1.1", "method":"GET", "host":"localhost:8080", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":200, "encoding":null, "bytes":0, "timeout":-1, "rate-limit":-1, "rate-burst":-1, "redirect":10, "cc":"RD" }

}
