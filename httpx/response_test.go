package httpx

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	testResponse = "file://[cwd]/httpxtest/test-response.txt"
)

func ExampleTransformBody() {
	err := TransformBody(nil)
	fmt.Printf("test: TransformBody() -> [err:%v]\n", err)

	err = TransformBody(&http.Response{})
	fmt.Printf("test: TransformBody() -> [err:%v]\n", err)

	resp := &http.Response{StatusCode: http.StatusGatewayTimeout, Body: emptyReader}
	err = TransformBody(resp)
	fmt.Printf("test: TransformBody() -> [err:%v]\n", err)
	buf, err1 := io.ReadAll(resp.Body)
	fmt.Printf("test: io.ReadAll() -> [buf:%v] [err:%v]\n", string(buf), err1)

	resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("this is content")))}
	err = TransformBody(resp)
	fmt.Printf("test: TransformBody() -> [err:%v]\n", err)
	buf, err1 = io.ReadAll(resp.Body)
	fmt.Printf("test: io.ReadAll() -> [buf:%v] [err:%v]\n", string(buf), err1)

	//Output:
	//test: TransformBody() -> [err:<nil>]
	//test: TransformBody() -> [err:<nil>]
	//test: TransformBody() -> [err:<nil>]
	//test: io.ReadAll() -> [buf:] [err:<nil>]
	//test: TransformBody() -> [err:<nil>]
	//test: io.ReadAll() -> [buf:this is content] [err:<nil>]

}

func readAll(body io.ReadCloser) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	defer body.Close()
	buf, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func ExampleNewResponse_Error() {
	resp, _ := NewResponse(http.StatusGatewayTimeout, nil, nil)
	buf, _ := iox.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	resp, _ = NewResponse(http.StatusGatewayTimeout, nil, errors.New("Deadline Exceeded"))
	buf, _ = iox.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:504] [content:]
	//test: NewResponse() -> [status-code:504] [content:Deadline Exceeded]

}

func ExampleNewResponse() {
	resp, _ := NewResponse(http.StatusOK, nil, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp, _ = NewResponse(http.StatusOK, nil, "version 1.2.35")
	buf, _ := iox.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:200]
	//test: NewResponse() -> [status-code:200] [content:version 1.2.35]

}
func ExampleNewHealthResponseOK() {
	status := "\"status\": \"up\""
	resp := NewHealthResponseOK()
	buf, _ := iox.ReadAll(resp.Body, nil)
	body := string(buf)
	fmt.Printf("test: NewHealthResponseOK() -> [status-code:%v] [content:%v]\n", resp.StatusCode, strings.Contains(body, status))

	//Output:
	//test: NewHealthResponseOK() -> [status-code:200] [content:true]

}

func ExampleNewNotFoundResponseWithStatus() {
	resp := NewNotFoundResponse()
	buf, _ := iox.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewNotFoundResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewNotFoundResponse() -> [status-code:404] [content:Not Found]

}

func Example_NewResponseFromUri() {
	s := testResponse
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [status:%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	buf, status := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [status:%v] [content-length:%v]\n", status, len(buf)) //string(buf))

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/test-response.txt) -> [status:<nil>] [statusCode:200]
	//test: readAll() -> [status:<nil>] [content-length:56]

}

func Example_NewResponseFromUri_URL_Nil() {
	resp, status0 := NewResponseFromUri(nil)
	fmt.Printf("test: NewResponseFromUri(nil) -> [%v] [statusCode:%v]\n", status0, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(nil) -> [error: URL is nil] [statusCode:500]

}

func _Example_NewResponseFromUri_Invalid_Scheme() {
	s := "https://www.google.com/search?q=golang"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%vl) -> [error:%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(https://www.google.com/search?q=golangl) -> [error:[error: Invalid URL scheme : https]] [statusCode:500]

}

func Example_NewResponseFromUri_HTTP_Error() {
	s := "file://[cwd]/httpxtest/message.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/message.txt) -> [malformed HTTP status code "text"] [statusCode:500]

}

func Example_NewResponseFromUri_504() {
	s := "file://[cwd]/httpxtest/http-504.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	buf, status := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [status:%v] [content-length:%v]\n", status, len(buf)) //string(buf))

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/http-504.txt) -> [<nil>] [statusCode:504]
	//test: readAll() -> [status:<nil>] [content-length:0]

}

func Example_NewResponseFromUri_EOF_Error() {
	s := "file://[cwd]/httpxtest/http-503-error.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/http-503-error.txt) -> [unexpected EOF] [statusCode:500]

}

/*
func ExampleNewError() {
	status := messaging.StatusOK()
	//var resp *httpx.Response

	err := NewError(nil, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, nil, err)

	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", messaging.StatusOK(), nil, err)

	status = messaging.NewStatusError(messaging.StatusInvalidContent, errors.New("error: invalid content"))
	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", status, nil, err)

	resp, _ := NewResponse(httpx.StatusTeapot, nil, nil)
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, resp != nil, err)

	resp, _ = NewResponse(httpx.StatusTeapot, nil, "error: response content")
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", nil, resp != nil, err)

	//Output:
	//test: NewError() -> [status:<nil>] [resp:<nil>] [err:]
	//test: NewError() -> [status:OK] [resp:<nil>] [err:]
	//test: NewError() -> [status:Invalid Content [error: invalid content]] [resp:<nil>] [error: invalid content]
	//test: NewError() -> [status:<nil>] [resp:true] [err:]
	//test: NewError() -> [status:<nil>] [resp:true] [error: response content]

}


*/
