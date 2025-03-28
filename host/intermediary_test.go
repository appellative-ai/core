package host

import (
	"context"
	"github.com/behavioral-ai/core/access"
	"github.com/behavioral-ai/core/httpx"
	"net/http"
	"time"
)

/*
func serviceTestExchange(_ *http.Request) (*http.Response, error) {
	//w.WriteHeader(httpx.StatusOK)
	//fmt.Fprint(w, "Service OK")
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("Service OK")))}, nil
}

func authTestExchange(r *http.Request) (*http.Response, error) {
	if r != nil {
		tokenString := r.Header.Get(Authorization)
		if tokenString == "" {
			//w.WriteHeader(httpx.StatusUnauthorized)
			//fmt.Fprint(w, "Missing authorization header")
			return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(bytes.NewReader([]byte("Missing authorization header")))}, errors.New("httpx.StatusUnauthorized")
		}
	}
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("Authorized")))}, nil
}

func ExampleConditionalIntermediary_Nil() {
	ic := NewConditionalIntermediary(nil, nil, nil)
	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)

	_, status := ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-auth -> [status:%v]\n", status)

	ic = NewConditionalIntermediary(authTestExchange, nil, nil)
	_, status = ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-service -> [status:%v]\n", status)

	//Output:
	//test: ConditionalIntermediary()-nil-auth -> [status:c1 is nil]
	//test: ConditionalIntermediary()-nil-service -> [status:c2 is nil]

}

func ExampleConditionalIntermediary_AuthExchange() {
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
	//test: ConditionalIntermediary()-auth-failure -> [status:httpx.StatusUnauthorized] [content:Missing authorization header]
	//test: ConditionalIntermediary()-auth-success -> [status:<nil>] [content:Service OK]

}

func _ExampleAccessLogIntermediary() {
	ic := NewAccessLogIntermediary(access.IngressTraffic, testDo)

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
	//req, _ := httpx.NewRequestWithContext(r.Context(), httpx.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
	}
	if resp == nil {
		resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
		return resp, errors.New("httpx.StatusGatewayTimeout")
	} else {
		resp.Body = io.NopCloser(bytes.NewReader([]byte("200 OK")))
		return resp, errors.New(fmt.Sprintf("status code %v", resp.StatusCode))
	}
}

func ExampleProxyIntermediary() {
	host := "www.search.yahoo.com"
	proxy := NewProxyIntermediary(host, proxyDo)

	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, status := proxy(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: ProxyIntermediary()-OK -> [status:%v] [content:%v]\n", status, len(buf) > 0)

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

*/

func limitLink(next httpx.Exchange) httpx.Exchange {
	return func(req *http.Request) (*http.Response, error) {
		time.Sleep(time.Second * 3)
		h := make(http.Header)
		h.Add(access.XRateLimit, "123")
		h.Add(access.XRateBurst, "12")
		return &http.Response{StatusCode: http.StatusTooManyRequests, Header: h}, nil
	}
}

/*
func timeoutExchange(next httpx.Exchange) httpx.Exchange {
	return func(req *http.Request) (*http.Response, error) {
		time.Sleep(time.Second * 3)
		h := make(http.Header)
		//h.Add(access.XRateLimit, "123")
		//h.Add(access.XRateBurst, "12")
		return &http.Response{StatusCode: http.StatusGatewayTimeout, Header: h}, nil
	}
}


*/

func ExampleAccessLogLink() {
	access.SetOrigin(access.Origin{
		Region:     "us-west1",
		Zone:       "zone-a",
		SubZone:    "sub-zone-1",
		Host:       "test.com",
		InstanceId: "123456789",
	})
	//ctx, fn := context.WithTimeout(context.Background(), time.Second*2)
	//defer fn()
	//req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/search?q=golang", nil)

	//ex := httpx.NewPipeline(AccessLogExchange, timeoutExchange)
	//ex(req)

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	req.Header.Add(access.XRequestId, "request-id")
	ex := httpx.BuildChain(AccessLogLink, limitLink)
	ex(req)

	//Output:
	//test: AccessLogIntermediary()-OK -> [status:<nil>] [content:true]
	//test: AccessLogIntermediary()-Gateway-Timeout -> [status:status code 504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}
