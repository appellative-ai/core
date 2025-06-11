package host

import (
	"context"
	access "github.com/behavioral-ai/core/access2"
	"github.com/behavioral-ai/core/rest"
	"net/http"
	"time"
)

func limitLink(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (*http.Response, error) {
		time.Sleep(time.Second * 3)
		h := make(http.Header)
		h.Add(access.XRateLimit, "123")
		return &http.Response{StatusCode: http.StatusTooManyRequests, Header: h}, nil
	}
}

func timeoutLink(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (*http.Response, error) {
		time.Sleep(time.Second * 3)
		h := make(http.Header)
		h.Add(access.XTimeout, "1500ms")
		return &http.Response{StatusCode: http.StatusGatewayTimeout, Header: h}, nil
	}
}

func redirectLink(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (*http.Response, error) {
		time.Sleep(time.Second * 3)
		h := make(http.Header)
		h.Add(access.XRedirect, "51")
		h.Add(access.XTimeout, "1500ms")
		return &http.Response{StatusCode: http.StatusGatewayTimeout, Header: h}, nil
	}
}

func ExampleAccessLogLink_Limit() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	req.Header.Add(access.XRequestId, "request-id")
	ex := rest.BuildExchangeChain([]any{accessLogLink, limitLink})
	ex(req)

	//Output:
	//test: AccessLogIntermediary()-OK -> [status:<nil>] [content:true]

}

func ExampleAccessLogLink_Timeout() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	req.Header.Add(access.XRequestId, "request-id")
	ex := rest.BuildExchangeChain([]any{accessLogLink, timeoutLink})
	ex(req)

	//Output:
	//test: AccessLogIntermediary()-OK -> [status:<nil>] [content:true]
	//test: AccessLogIntermediary()-Gateway-Timeout -> [status:status code 504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func ExampleAccessLogLink_Redirect() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	req.Header.Add(access.XRequestId, "request-id")
	ex := rest.BuildExchangeChain([]any{accessLogLink, redirectLink})
	ex(req)

	//Output:
	//test: AccessLogIntermediary()-OK -> [status:<nil>] [content:true]

}
