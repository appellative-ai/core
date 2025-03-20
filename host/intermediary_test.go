package host

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/access"
	"io"
	"net/http"
	"time"
)

func serviceTestExchange(_ *http.Request) (*http.Response, error) {
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, "Service OK")
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("Service OK")))}, nil
}

func authTestExchange(r *http.Request) (*http.Response, error) {
	if r != nil {
		tokenString := r.Header.Get(Authorization)
		if tokenString == "" {
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprint(w, "Missing authorization header")
			return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(bytes.NewReader([]byte("Missing authorization header")))}, errors.New("http.StatusUnauthorized")
		}
	}
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("Authorized")))}, nil
}

func _ExampleConditionalIntermediary_Nil() {
	ic := NewConditionalIntermediary(nil, nil, nil)
	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)

	_, status := ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-auth -> [status:%v]\n", status)

	ic = NewConditionalIntermediary(authTestExchange, nil, nil)
	_, status = ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-service -> [status:%v]\n", status)

	//Output:
	//test: ConditionalIntermediary()-nil-auth -> [status:error: Conditional Intermediary HttpExchange 1 is nil]
	//test: ConditionalIntermediary()-nil-service -> [status:error: Conditional Intermediary HttpExchange 2 is nil]

}

func _ExampleConditionalIntermediary_AuthExchange() {
	ic := NewConditionalIntermediary(authTestExchange, serviceTestExchange, nil)
	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)

	resp, status := ic(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: ConditionalIntermediary()-auth-failure -> [status:%v] [content:%v]\n", status, string(buf))

	r.Header.Add(Authorization, "token")
	resp, status = ic(r)
	buf, _ = io.ReadAll(resp.Body)
	fmt.Printf("test: ConditionalIntermediary()-auth-success -> [status:%v] [content:%v]\n", status, string(buf))

	//Output:
	//test: ConditionalIntermediary()-auth-failure -> [status:http.StatusUnauthorized] [content:Missing authorization header]
	//test: ConditionalIntermediary()-auth-success -> [status:<nil>] [content:Service OK]

}

func _ExampleAccessLogIntermediary() {
	ic := NewAccessLogIntermediary(access.InternalTraffic, testDo)

	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)
	resp, status := ic(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: AccessLogIntermediary()-OK -> [status:%v] [content:%v]\n", status, len(buf) > 0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()
	r, _ = http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/search?q-golang", nil)
	resp, status = ic(r)
	buf = nil
	if resp.Body != nil {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: AccessLogIntermediary()-Gateway-Timeout -> [status:%v] [content:%v]\n", status, string(buf))

	//Output:
	//test: AccessLogIntermediary()-OK -> [status:<nil>] [content:true]
	//test: AccessLogIntermediary()-Gateway-Timeout -> [status:status code 504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func proxyDo(r *http.Request) (*http.Response, error) {
	//req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
	}
	if resp == nil {
		resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
		return resp, errors.New("http.StatusGatewayTimeout")
	} else {
		resp.Body = io.NopCloser(bytes.NewReader([]byte("200 OK")))
		return resp, errors.New(fmt.Sprintf("status code %v", resp.StatusCode))
	}
}

func _ExampleProxyIntermediary() {
	host := "www.search.yahoo.com"
	proxy := NewProxyIntermediary(host, proxyDo)

	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, status := proxy(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: ProxyIntermediary()-OK -> [status:%v] [content:%v]\n", status, len(buf) > 0)

	/*
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
		defer cancel()
		r, _ = http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/search?q-golang", nil)
		resp, status = ic(r)
		buf = nil
		if resp.Body != nil {
			buf, _ = io.ReadAll(resp.Body)
		}
		fmt.Printf("test: AccessLogIntermediary()-Gateway-Timeout -> [status:%v] [content:%v]\n", status, string(buf))


	*/

	//Output:
	//test: ProxyIntermediary()-OK -> [status:status code 500] [content:true]

}

func testDo(r *http.Request) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	if resp == nil {
		resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
		return resp, errors.New(fmt.Sprintf("status code %v", resp.StatusCode))
	} else {
		resp.Body = io.NopCloser(bytes.NewReader([]byte("200 OK")))
		return resp, nil //errors.New(fmt.Sprintf("status code %v",resp.StatusCode))
	}
}

type node struct {
	ex func(req *http.Request, next *node) (*http.Response, error)
	n2 func() *node
}

func do1(req *http.Request, next *node) (*http.Response, error) {
	fmt.Printf("test: do1() -> request\n")
	if next != nil {
		next.ex(req, next.n2())

	}
	fmt.Printf("test: do1() -> response\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func do2(req *http.Request, next *node) (*http.Response, error) {
	fmt.Printf("test: do2() -> request\n")
	if next != nil {
		next.ex(req, next.n2())
	}
	fmt.Printf("test: do2() -> response\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func do3(req *http.Request, next *node) (*http.Response, error) {
	fmt.Printf("test: do3() -> request\n")
	if next != nil {
		next.ex(req, next.n2())
	}
	fmt.Printf("test: do3() -> response\n")
	return &http.Response{StatusCode: http.StatusBadRequest}, nil
}

func do4(req *http.Request) (*http.Response, error) {
	fmt.Printf("test: do4()\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func do5(req *http.Request) (*http.Response, error) {
	fmt.Printf("test: do5()\n")
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func ExampleLinkedExchange() {
	//fn := func(code int) bool { return code == http.StatusOK }
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)

	do1(req, &node{ex: do2, n2: func() *node { return &node{ex: do3, n2: func() *node { return nil }} }})
	//}do2)
	//do2(req, nil)

	//Output:
	//fail
}

func _ExampleIntermediary_Cond() {
	//fn := func(code int) bool { return code == http.StatusOK }
	//req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)

	//ex := NewConditionalIntermediary(do1, do2, fn)
	//ex = NewConditionalIntermediary(ex, do3, fn)
	//ex = NewConditionalIntermediary(ex, do4, fn)
	//ex = NewConditionalIntermediary(ex, do5, fn)

	//ex(req)

	//Output:
	//fail
}
