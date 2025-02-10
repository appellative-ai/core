package aspect

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var testTS time.Time

func init() {
	testTS = time.Date(2024, 3, 1, 18, 23, 50, 205*1e6, time.UTC)

}

func Example_FormatUri() {
	s := "github/behavioral-ai/core/core:testFunc"

	fmt.Printf("test: formatUri(%v) -> %v\n", s, formatUri(s))

	s = "gitlab/behavioral-ai/core/core:testFunc"
	fmt.Printf("test: formatUri(%v) -> %v\n", s, formatUri(s))

	//Output:
	//test: formatUri(github/behavioral-ai/core/core:testFunc) -> https://github.com/behavioral-ai/core/tree/main/core#testFunc
	//test: formatUri(gitlab/behavioral-ai/core/core:testFunc) -> gitlab/behavioral-ai/core/core:testFunc

}

func Example_FormatUri_Test() {
	s := "http://localhost:8080/github.com/behavioral-ai/core/core/testFunc"
	req, err := http.NewRequest("", s, nil)
	fmt.Printf("test: http.URL -> [req:%v] [url:%v] [err:%v]\n", req != nil, req.URL, err)

	s = "http://localhost:8080/github.com/behavioral-ai/core/core:testFunc"
	req, err = http.NewRequest("", s, nil)
	fmt.Printf("test: http.URL -> [req:%v] [url:%v] [err:%v]\n", req != nil, req.URL, err)

	s = "http://localhost:8080/github.com:behavioral-ai/core/core.testFunc"
	req, err = http.NewRequest("", s, nil)
	fmt.Printf("test: http.URL -> [req:%v] [url:%v] [err:%v]\n", req != nil, req.URL, err)

	//Output:
	//test: http.URL -> [req:true] [url:http://localhost:8080/github.com/behavioral-ai/core/core/testFunc] [err:<nil>]
	//test: http.URL -> [req:true] [url:http://localhost:8080/github.com/behavioral-ai/core/core:testFunc] [err:<nil>]
	//test: http.URL -> [req:true] [url:http://localhost:8080/github.com:behavioral-ai/core/core.testFunc] [err:<nil>]

}

func Example_DefaultFormat() {
	s := NewStatusError(http.StatusNotFound, errors.New("test error message 1"))

	str := formatter(testTS, s.Code, HttpStatus(s.Code), "1234-5678", []error{s.Err}, s.Trace())
	fmt.Printf("test: formatter() -> %v", str)

	//Output:
	//test: formatter() -> { "timestamp":"2024-03-01T18:23:50.205Z", "code":404, "status":"Not Found", "request-id":"1234-5678", "errors" : [ "test error message 1" ], "trace" : [ "https://github.com/behavioral-ai/core/tree/main/core#Example_DefaultFormat" ] }

}
