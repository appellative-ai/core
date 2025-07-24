package httpx

import (
	"bytes"
	"errors"
	"fmt"
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
	fmt.Printf("test: TransformBody() -> [cnt:%v] [err:%v]\n", 0, err)

	err = TransformBody(&http.Response{})
	fmt.Printf("test: TransformBody() -> [cnt:%v] [err:%v]\n", 0, err)

	resp := &http.Response{StatusCode: http.StatusGatewayTimeout, Body: EmptyReader}
	err = TransformBody(resp)
	fmt.Printf("test: TransformBody() -> [cnt:%v] [err:%v]\n", resp.ContentLength, err)
	buf, err1 := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [buf:%v] [err:%v]\n", string(buf), err1)

	resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("this is content")))}
	err = TransformBody(resp)
	fmt.Printf("test: TransformBody() -> [cnt:%v] [err:%v]\n", resp.ContentLength, err)
	buf, err1 = readAll(resp.Body)
	fmt.Printf("test: readAll() -> [buf:%v] [err:%v]\n", string(buf), err1)

	//Output:
	//test: TransformBody() -> [cnt:0] [err:<nil>]
	//test: TransformBody() -> [cnt:0] [err:<nil>]
	//test: TransformBody() -> [cnt:0] [err:<nil>]
	//test: readAll() -> [buf:] [err:<nil>]
	//test: TransformBody() -> [cnt:15] [err:<nil>]
	//test: readAll() -> [buf:this is content] [err:<nil>]

}

func readAll2(body io.ReadCloser) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	defer body.Close()
	buf, err := readAll(body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func ExampleNewResponse_Error() {
	resp := NewResponse(http.StatusGatewayTimeout, nil, nil)
	buf, _ := readAll(resp.Body)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	resp = NewResponse(http.StatusGatewayTimeout, nil, errors.New("Deadline Exceeded"))
	buf, _ = readAll(resp.Body)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:504] [content:]
	//test: NewResponse() -> [status-code:504] [content:Deadline Exceeded]

}

func ExampleNewResponse() {
	resp := NewResponse(http.StatusOK, nil, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp = NewResponse(http.StatusOK, nil, "version 1.2.35")
	buf, _ := readAll(resp.Body)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:200]
	//test: NewResponse() -> [status-code:200] [content:version 1.2.35]

}
func ExampleNewHealthResponseOK() {
	status := "\"status\": \"up\""
	resp := NewHealthResponseOK()
	buf, _ := readAll(resp.Body)
	body := string(buf)
	fmt.Printf("test: NewHealthResponseOK() -> [status-code:%v] [content:%v]\n", resp.StatusCode, strings.Contains(body, status))

	//Output:
	//test: NewHealthResponseOK() -> [status-code:200] [content:true]

}

func ExampleNewNotFoundResponseWithStatus() {
	resp := NewNotFoundResponse()
	buf, _ := readAll(resp.Body)
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
	status := std.StatusOK
	//var resp *httpx.Response

	err := NewError(nil, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, nil, err)

	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", std.StatusOK, nil, err)

	status = std.NewStatus(std.StatusInvalidContent, "",errors.New("error: invalid content"))
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
