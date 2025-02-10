package httpx

import (
	"fmt"
	"io"
	"net/http"
)

func Example_ResponseWriter() {
	requestId2 := "123-request-id"
	relatesTo2 := "test-relates-to"
	content := "this is response write content"
	w := NewResponseWriter()

	w.Header().Add(aspect.XRequestId, requestId2)
	w.Header().Add(aspect.XRelatesTo, relatesTo2)
	w.WriteHeader(http.StatusAccepted)
	cnt, err := w.Write([]byte(content))
	fmt.Printf("test: responseWriter() -> [cnt:%v] [error:%v]\n", cnt, err)

	resp := w.Response()

	fmt.Printf("test: responseWriter() -> [write-requestId:%v] [response-requestId:%v]\n", requestId, resp.Header.Get(aspect.XRequestId))
	fmt.Printf("test: responseWriter() -> [write-relatesTo:%v] [response-relatesTo:%v]\n", relatesTo, resp.Header.Get(aspect.XRelatesTo))
	fmt.Printf("test: responseWriter() -> [write-statusCode:%v] [response-statusCode:%v]\n", http.StatusAccepted, resp.StatusCode)

	buf, _ := io.ReadAll(resp.Body)

	fmt.Printf("test: responseWriter() -> [write-content:%v] [response-content:%v]\n", content, string(buf))

	//Output:
	//test: responseWriter() -> [cnt:30] [error:<nil>]
	//test: responseWriter() -> [write-requestId:123-request-id] [response-requestId:123-request-id]
	//test: responseWriter() -> [write-relatesTo:test-relates-to] [response-relatesTo:test-relates-to]
	//test: responseWriter() -> [write-statusCode:202] [response-statusCode:202]
	//test: responseWriter() -> [write-content:this is response write content] [response-content:this is response write content]

}
