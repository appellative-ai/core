package host

import (
	"bytes"
	"fmt"
	"github.com/behavioral-ai/core/access"
	"github.com/behavioral-ai/core/aspect"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func appHttpExchange(r *http.Request) (*http.Response, *aspect.Status) {
	status := aspect.NewStatus(http.StatusTeapot)
	return &http.Response{StatusCode: status.Code}, status
}

func testAuthExchangeOK(_ *http.Request) (*http.Response, *aspect.Status) {
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, "OK")
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("200 OK")))}, aspect.StatusOK()
}

func testAuthExchangeFail(_ *http.Request) (*http.Response, *aspect.Status) {
	//w.WriteHeader(http.StatusUnauthorized)
	//fmt.Fprint(w, "Missing authorization header")
	return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(bytes.NewReader([]byte("Missing authorization header")))}, aspect.NewStatus(http.StatusUnauthorized)
}

func testDo(r *http.Request) (*http.Response, *aspect.Status) {
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	if resp == nil {
		resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
		return resp, aspect.NewStatus(http.StatusGatewayTimeout)
	} else {
		resp.Body = io.NopCloser(bytes.NewReader([]byte("200 OK")))
		return resp, aspect.NewStatus(resp.StatusCode)
	}
}

/*
	func ExampleRegisterdomain() {
		domain := []PathHandler{
			{"test", nil},
		}
		err := Registerdomain(domain)
		fmt.Printf("test: Registerdomain() -> [%v]\n", err)

		domain = []PathHandler{
			{"test", testAuthExchangeOK},
			{"test2", testAuthExchangeOK},
		}
		err = Registerdomain(domain)
		fmt.Printf("test: Registerdomain() -> [%v]\n", err)

		//Output:
		//test: Registerdomain() -> [error: handler for path [test] is nil]
		//test: Registerdomain() -> [<nil>]

}
*/

func ExampleRegisterExchange() {
	a := "github/advanced-go/stdlib"

	err := RegisterExchange("", nil)
	fmt.Printf("test: RegisterExchange(_,nil) -> [err:%v]\n", err)

	err = RegisterExchange(a, nil)
	fmt.Printf("test: RegisterExchange(\"%v\",nil) -> [err:%v]\n", a, err)

	err = RegisterExchange(a, appHttpExchange)
	fmt.Printf("test: RegisterExchange(\"%v\",appHttpExchange) -> [err:%v]\n", a, err)

	h := exchangeProxy.Lookup(a)
	fmt.Printf("test: Lookup(\"%v\") -> [ok:%v]\n", a, h != nil)

	err = RegisterExchange(a, appHttpExchange)
	fmt.Printf("test: RegisterExchange(\"%v\",appHttpExchange) -> [err:%v]\n", a, err)

	//Output:
	//test: RegisterExchange(_,nil) -> [err:invalid argument: domain is empty]
	//test: RegisterExchange("github/advanced-go/stdlib",nil) -> [err:invalid argument: HTTP Exchange is nil for domain : [github/advanced-go/stdlib]]
	//test: RegisterExchange("github/advanced-go/stdlib",appHttpExchange) -> [err:<nil>]
	//test: Lookup("github/advanced-go/stdlib") -> [ok:true]
	//test: RegisterExchange("github/advanced-go/stdlib",appHttpExchange) -> [err:invalid argument: HTTP Exchange already exists for domain : [github/advanced-go/stdlib]]

}

func ExampleHttpHandler() {
	pattern := "github/advanced-go/host"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host:entry", nil)

	RegisterExchange(pattern, appHttpExchange)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)

	fmt.Printf("test: HttpHandler() -> %v\n", rec.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> 418

}

func ExampleHttpHandler_Host_OK() {
	pattern := "github/advanced-go/host/ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/ok:entry", nil)

	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_Host_Timeout() {
	pattern := "github/advanced-go/host/timeout"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/timeout:entry", nil)

	SetHostTimeout(time.Millisecond)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func ExampleHttpHandler_Auth_Authorized() {
	pattern := "github/advanced-go/host/authorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/authorized:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_Auth_Unauthorized() {
	pattern := "github/advanced-go/host/unauthorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/unauthorized:entry", nil)

	SetAuthExchange(testAuthExchangeFail, nil)
	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:401] [content:Missing authorization header]

}

func ExampleHttpHandler_AccessLog_Service_OK() {
	pattern := "github/advanced-go/host/access-log-service-ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/access-log-service-ok:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout(time.Second * 4)
	RegisterExchange(pattern, NewAccessLogIntermediary(access.InternalTraffic, testDo))

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_AccessLog_Service_Timeout() {
	pattern := "github/advanced-go/host/access-log-service-timeout"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/access-log-service-timeout:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout(time.Millisecond * 4)
	RegisterExchange(pattern, NewAccessLogIntermediary(access.InternalTraffic, testDo))

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func ExampleHttpHandler_AccessLog_Service_Unauthorized() {
	pattern := "github/advanced-go/host/access-log-service-unauthorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/access-log-service-unauthorized:entry", nil)

	SetAuthExchange(testAuthExchangeFail, nil)
	SetHostTimeout(time.Second * 4)
	RegisterExchange(pattern, NewAccessLogIntermediary(access.InternalTraffic, testDo))

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:401] [content:Missing authorization header]

}
