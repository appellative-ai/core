package test

import (
	"fmt"
	"github.com/behavioral-ai/core/httpx"
)

const (
	testInputResponse  = "file://[cwd]/test/test-response.txt"
	testOutputResponse = "file://[cwd]/test/test-response-2.txt"
)

func _ExampleWriteStatusLine() {
	resp, status := httpx.NewResponseFromUri(testInputResponse)
	line := fmt.Sprintf("%v %v\n", resp.Proto, resp.Status)
	fmt.Printf("test: NewResponseFromUri() [status:%v] [%v]\n", status, line)

	//Output:
	//fail

}

func ExampleWriteResponse() {
	resp, status := httpx.NewResponseFromUri(testInputResponse)
	fmt.Printf("test: NewResponseFromUri() [status:%v] [%v]\n", status, resp.Proto+" "+resp.Status)

	status = WriteResponse(testOutputResponse, resp)
	fmt.Printf("test: WriteResponse(\"%v\") -> [status:%v]\n", testInputResponse, status)

	//Output:
	//test: NewResponseFromUri() [status:OK] [HTTP/1.1 200 OK]
	//test: WriteResponse("file://[cwd]/test/test-response.txt") -> [status:OK]

}
