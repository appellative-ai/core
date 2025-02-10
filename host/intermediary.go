package host

import (
	"errors"
	"github.com/behavioral-ai/core/access"
	"github.com/behavioral-ai/core/httpx"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
)

func badRequest(msg string) (*http.Response, *aspect.Status) {
	return &http.Response{StatusCode: http.StatusBadRequest}, aspect.NewStatusError(http.StatusBadRequest, errors.New(msg))
}

func NewConditionalIntermediary(c1 aspect.HttpExchange, c2 aspect.HttpExchange, ok func(int) bool) aspect.HttpExchange {
	return func(r *http.Request) (resp *http.Response, status *aspect.Status) {
		if c1 == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 1 is nil")
		}
		if c2 == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 2 is nil")
		}
		resp, status = c1(r)
		if resp == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 1 response is nil")
		}
		if (ok == nil && resp.StatusCode == http.StatusOK) || (ok != nil && ok(resp.StatusCode)) {
			resp, status = c2(r)
		}
		return
	}
}

func NewAccessLogIntermediary(traffic string, c2 aspect.HttpExchange) aspect.HttpExchange {
	return func(r *http.Request) (resp *http.Response, status *aspect.Status) {
		if c2 == nil {
			return badRequest("error: AccessLog Intermediary HttpExchange is nil")
		}
		reasonCode := ""
		from := r.Header.Get(aspect.XFrom)

		var dur time.Duration
		if ct, ok := r.Context().Deadline(); ok {
			dur = time.Until(ct) * -1
		}
		start := time.Now().UTC()
		resp, status = c2(r)
		if status.Code == http.StatusGatewayTimeout {
			reasonCode = access.ControllerTimeout
		}
		route := resp.Header.Get(aspect.XRoute)
		if route == "" {
			route = EtcRoute
		}
		access.Log(traffic, start, time.Since(start), r, resp, access.Routing{From: from, Route: route, To: "", Percent: -1}, access.Controller{Timeout: dur, RateLimit: 0, RateBurst: 0, Code: reasonCode})
		return
	}
}

func NewProxyIntermediary(host string, c2 aspect.HttpExchange) aspect.HttpExchange {
	return func(r *http.Request) (resp *http.Response, status *aspect.Status) {
		if c2 == nil {
			return badRequest("error: Proxy Intermediary HttpExchange is nil")
		}
		u := r.URL.Scheme + "://" + host + r.URL.Path
		if r.URL.RawQuery != "" {
			u += "?" + r.URL.RawQuery
		}
		r2, err := http.NewRequestWithContext(r.Context(), r.Method, u, r.Body)
		if err != nil {
			return httpx.NewResponse(http.StatusBadRequest, r.Header, err)
		}
		return c2(r2)
	}
}
