package core

import "net/http"

const (
	XRequestId        = "X-Request-Id"
	XRelatesTo        = "X-Relates-To"
	XDomain           = "X-Domain"
	XVersion          = "X-Version"
	XURLPath          = "x-url-path"
	XTest             = "X-Test"
	XFrom             = "X-From"
	XTo               = "X-To"
	XRoute            = "X-Route"
	XExchangeRequest  = "X-Exchange-Request"
	XExchangeResponse = "X-Exchange-Response"
	XExchangeStatus   = "X-Exchange-Status"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)

func ExchangeHeaders(h http.Header) (req, resp, status string) {
	if h == nil {
		return
	}
	return h.Get(XExchangeRequest), h.Get(XExchangeResponse), h.Get(XExchangeStatus)
}
