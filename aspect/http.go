package aspect

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

	IAMUser           = "X-User"
	IAMPassword       = "X-Password"
	IAMUri            = "X-Uri"
	IAMFrom           = "X-From"
	IAMCredentialsUri = "iam:credentials"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)

func ExchangeHeaders(h http.Header) (req, resp, status string) {
	if h == nil {
		return
	}
	return h.Get(XExchangeRequest), h.Get(XExchangeResponse), h.Get(XExchangeStatus)
}

// Resource - resource interface
type Resource interface {
	HttpExchange
}

func NewIAMRequest(from string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, IAMCredentialsUri, nil)
	req.Header = make(http.Header)
	req.Header.Set(IAMFrom, from)
	return req
}

func NewIAMResponseFromUri(uri string, status *Status) *http.Response {
	resp := &http.Response{Status: status.String(), StatusCode: status.HttpCode(), Header: make(http.Header)}
	resp.Header.Set(IAMUri, uri)
	return resp
}

func NewIAMResponseFromCredentials(user, password string, status *Status) *http.Response {
	resp := &http.Response{Status: status.String(), StatusCode: status.HttpCode(), Header: make(http.Header)}
	resp.Header.Set(IAMPassword, password)
	resp.Header.Set(IAMUser, user)
	return resp
}
