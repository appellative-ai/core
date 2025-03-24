package host

import (
	"github.com/behavioral-ai/core/access"
	"github.com/behavioral-ai/core/httpx"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
)

/*
func badRequest(msg string) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusBadRequest}, errors.New(msg)
}

func NewConditionalIntermediary(c1 httpx.Exchange, c2 httpx.Exchange, ok func(int) bool) httpx.Exchange {
	if ok == nil {
		ok = func(code int) bool { return code == http.StatusOK }
	}
	return func(r *http.Request) (resp *http.Response, err error) {
		if c1 == nil {
			return badRequest("c1 is nil")
		}
		if c2 == nil {
			return badRequest("c2 is nil")
		}
		resp, err = c1(r)
		if resp == nil {
			return &http.Response{StatusCode: http.StatusBadRequest}, err
		}
		if ok(resp.StatusCode) {
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
		//from := r.Header.Get(XFrom)

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
		access.Log(traffic, start, time.Since(start), r, resp, access.Controller{Timeout: dur, RateLimit: "0", RateBurst: "0", Code: reasonCode})
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


*/

func AuthorizationExchange(next httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(Authorization)
		if auth != "" {
			return &http.Response{StatusCode: http.StatusOK}, nil
		}
		return &http.Response{StatusCode: http.StatusUnauthorized}, nil
	}
}

func AccessLogExchange(next httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		limit := ""
		burst := ""
		pct := ""

		/*
			var dur time.Duration
			if ct, ok := r.Context().Deadline(); ok {
				dur = time.Until(ct) * -1
			}
		*/

		if next != nil {
			resp, err = next(r)
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			limit = resp.Header.Get(access.XRateLimit)
			resp.Header.Del(access.XRateLimit)
			burst = resp.Header.Get(access.XRateBurst)
			resp.Header.Del(access.XRateBurst)
		}
		pct = resp.Header.Get(access.XRedirect)
		if pct != "" {
			resp.Header.Del(access.XRedirect)
		}
		access.Log(access.IngressTraffic, start, time.Since(start), r, resp, access.Controller{Timeout: -1, RateLimit: limit, RateBurst: burst, Redirect: pct})
		return
	}
}
