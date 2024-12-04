package httpx

import (
	"github.com/behavioral-ai/core/core"
	"net/http"
)

var (
	exchangeProxy = core.NewExchangeProxy()
)

// RegisterExchange - add a domain and Http Exchange handler to the proxy
func RegisterExchange(domain string, handler core.HttpExchange) error {
	return exchangeProxy.Register(domain, handler)
}

// Exchange - process an HTTP call utilizing an Exchange
func Exchange(req *http.Request) (*http.Response, *core.Status) {
	ex := exchangeProxy.LookupByRequest(req)
	if ex != nil {
		return ex(req)
	}
	//ctrl, status := controller2.Lookup(req)
	//if status.OK() {
	//	return controller2.Exchange(req, Do, ctrl)
	//}
	return Do(req)
}
