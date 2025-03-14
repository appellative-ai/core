package http

import (
	"errors"
	"fmt"
	io2 "github.com/behavioral-ai/core/io"
	"github.com/behavioral-ai/core/messaging"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	testResponse = "file://[cwd]/httptest/test-response.txt"
)

func readAll(body io.ReadCloser) ([]byte, *messaging.Status) {
	if body == nil {
		return nil, messaging.StatusOK()
	}
	defer body.Close()
	buf, err := io.ReadAll(body)
	if err != nil {
		return nil, messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return buf, messaging.StatusOK()
}

func ExampleNewResponse_Error() {
	status := messaging.NewStatus(http.StatusGatewayTimeout)
	resp, _ := NewResponse(status.HttpCode(), nil, status.Err)
	buf, _ := io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	status = messaging.NewStatusError(http.StatusGatewayTimeout, errors.New("Deadline Exceeded"), "")
	resp, _ = NewResponse(status.HttpCode(), nil, status.Err)
	buf, _ = io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:504] [content:]
	//test: NewResponse() -> [status-code:504] [content:Deadline Exceeded]

}

func ExampleNewResponse() {
	resp, _ := NewResponse(http.StatusOK, nil, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp, _ = NewResponse(messaging.StatusOK().HttpCode(), nil, "version 1.2.35")
	buf, _ := io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:200]
	//test: NewResponse() -> [status-code:200] [content:version 1.2.35]

}
func ExampleNewHealthResponseOK() {
	status := "\"status\": \"up\""
	resp := NewHealthResponseOK()
	buf, _ := io2.ReadAll(resp.Body, nil)
	body := string(buf)
	fmt.Printf("test: NewHealthResponseOK() -> [status-code:%v] [content:%v]\n", resp.StatusCode, strings.Contains(body, status))

	//Output:
	//test: NewHealthResponseOK() -> [status-code:200] [content:true]

}

func ExampleNewNotFoundResponseWithStatus() {
	resp := NewNotFoundResponse()
	buf, _ := io2.ReadAll(resp.Body, nil)
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
	//test: NewResponseFromUri(file://[cwd]/httptest/test-response.txt) -> [status:OK] [statusCode:200]
	//test: readAll() -> [status:OK] [content-length:56]

}

func Example_NewResponseFromUri_URL_Nil() {
	resp, status0 := NewResponseFromUri(nil)
	fmt.Printf("test: NewResponseFromUri(nil) -> [error:[%v]] [statusCode:%v]\n", status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(nil) -> [error:[error: URL is nil]] [statusCode:500]

}

func _Example_NewResponseFromUri_Invalid_Scheme() {
	s := "https://www.google.com/search?q=golang"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%vl) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(https://www.google.com/search?q=golangl) -> [error:[error: Invalid URL scheme : https]] [statusCode:500]

}

func Example_NewResponseFromUri_HTTP_Error() {
	s := "file://[cwd]/httptest/message.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httptest/message.txt) -> [error:[malformed HTTP status code "text"]] [statusCode:500]

}

func Example_NewResponseFromUri_504() {
	s := "file://[cwd]/httptest/http-504.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	buf, status := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [status:%v] [content-length:%v]\n", status, len(buf)) //string(buf))

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httptest/http-504.txt) -> [error:[<nil>]] [statusCode:504]
	//test: readAll() -> [status:OK] [content-length:0]

}

func Example_NewResponseFromUri_EOF_Error() {
	s := "file://[cwd]/httptest/http-503-error.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httptest/http-503-error.txt) -> [error:[unexpected EOF]] [statusCode:500]

}

/*
func ExampleNewError() {
	status := messaging.StatusOK()
	//var resp *http.Response

	err := NewError(nil, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, nil, err)

	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", messaging.StatusOK(), nil, err)

	status = messaging.NewStatusError(messaging.StatusInvalidContent, errors.New("error: invalid content"))
	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", status, nil, err)

	resp, _ := NewResponse(http.StatusTeapot, nil, nil)
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, resp != nil, err)

	resp, _ = NewResponse(http.StatusTeapot, nil, "error: response content")
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
