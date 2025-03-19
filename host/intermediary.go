package host

import (
	"errors"
	"github.com/behavioral-ai/core/access"
	httpx "github.com/behavioral-ai/core/http"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
)

func badRequest(msg string) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusBadRequest}, errors.New(msg)
}

func NewConditionalIntermediary(c1 httpx.Exchange, c2 httpx.Exchange, ok func(int) bool) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		if c1 == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 1 is nil")
		}
		if c2 == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 2 is nil")
		}
		resp, err = c1(r)
		if resp == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 1 response is nil")
		}
		if (ok == nil && resp.StatusCode == http.StatusOK) || (ok != nil && ok(resp.StatusCode)) {
			resp, err = c2(r)
		}
		return
	}
}

func NewAccessLogIntermediary(traffic string, c2 httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		if c2 == nil {
			return badRequest("error: AccessLog Intermediary HttpExchange is nil")
		}
		reasonCode := ""
		from := r.Header.Get(XFrom)

		var dur time.Duration
		if ct, ok := r.Context().Deadline(); ok {
			dur = time.Until(ct) * -1
		}
		start := time.Now().UTC()
		resp, err = c2(r)
		if resp.StatusCode == http.StatusGatewayTimeout {
			reasonCode = access.ControllerTimeout
		}
		route := resp.Header.Get(XRoute)
		if route == "" {
			route = EtcRoute
		}
		access.Log(traffic, start, time.Since(start), r, resp, access.Routing{From: from, Route: route, To: "", Percent: -1}, access.Controller{Timeout: dur, RateLimit: 0, RateBurst: 0, Code: reasonCode})
		return
	}
}

func NewProxyIntermediary(host string, c2 httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		if c2 == nil {
			return badRequest("error: Proxy Intermediary HttpExchange is nil")
		}
		u := r.URL.Scheme + "://" + host + r.URL.Path
		if r.URL.RawQuery != "" {
			u += "?" + r.URL.RawQuery
		}
		r2, err1 := http.NewRequestWithContext(r.Context(), r.Method, u, r.Body)
		if err1 != nil {
			resp, _ = httpx.NewResponse(http.StatusBadRequest, r.Header, err)
			return resp, nil //messaging.StatusOK()
		}
		return c2(r2)
	}
}

/*
func NewIntermediary(e1 httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		if e1 == nil {
			return badRequest("error: intermediary Exchange is nil")
		}
		resp,err := e1(r)


		return c2(r2)
	}
}


*/
