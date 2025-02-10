package httpx

import (
	"net/http"
)

var (
	exchangeProxy = aspect.NewExchangeProxy()
)

// RegisterExchange - add a domain and Http Exchange handler to the proxy
func RegisterExchange(domain string, handler aspect.HttpExchange) error {
	return exchangeProxy.Register(domain, handler)
}

// Exchange - process an HTTP call utilizing an Exchange
func Exchange(req *http.Request) (*http.Response, *aspect.Status) {
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
