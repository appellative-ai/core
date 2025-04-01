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

func AuthorizationLink(next httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(Authorization)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		if next != nil {
			resp, err = next(r)
		} else {
			return &http.Response{StatusCode: http.StatusOK}, nil
		}
		return
	}
}

func AccessLogLink(next httpx.Exchange) httpx.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		limit := ""
		pct := ""
		timeout := ""

		if next != nil {
			resp, err = next(r)
		}
		limit = resp.Header.Get(access.XRateLimit)
		resp.Header.Del(access.XRateLimit)
		timeout = resp.Header.Get(access.XTimeout)
		resp.Header.Del(access.XTimeout)
		pct = resp.Header.Get(access.XRedirect)
		resp.Header.Del(access.XRedirect)
		access.Log(access.IngressTraffic, start, time.Since(start), "", r, resp, access.Threshold{Timeout: timeout, RateLimit: limit, Redirect: pct})
		return
	}
}
