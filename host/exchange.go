package host

import (
	"github.com/behavioral-ai/core/uri"
	"net/http"
	"time"
)

var (
	exchangeProxy = aspect.NewExchangeProxy()
	hostDuration  time.Duration
	authExchange  aspect.HttpExchange
	okFunc        = func(code int) bool { return code == http.StatusOK }
)

func SetHostTimeout(d time.Duration) {
	hostDuration = d
}

func SetAuthExchange(h aspect.HttpExchange, ok func(int) bool) {
	if h != nil {
		authExchange = h
		if ok != nil {
			okFunc = ok
		}
	}
}

// RegisterExchange - add a domain and Http Exchange handler to the proxy
func RegisterExchange(domain string, handler aspect.HttpExchange) error {
	h := handler
	if authExchange != nil {
		h = NewConditionalIntermediary(authExchange, handler, okFunc)
	}
	return exchangeProxy.Register(domain, h)
}

// HttpHandler - process an HTTP request
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil || w == nil || r.URL == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := uri.Uproot(r.URL.Path)
	if !p.Valid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	handler := exchangeProxy.Lookup(p.Domain)
	if handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hostExchange(w, r, hostDuration, handler)
}
