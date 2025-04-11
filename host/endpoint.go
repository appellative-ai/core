package host

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"net/http"
)

// ExchangeHandler - http exchange handler interface
type ExchangeHandler2 interface {
	Exchange(w http.ResponseWriter, r *http.Request)
}

type endpoint struct {
	handler rest.Exchange
}

func NewEndpoint(handler rest.Exchange) ExchangeHandler2 {
	e := new(endpoint)
	e.handler = handler
	return e
}

func (e *endpoint) Exchange(w http.ResponseWriter, r *http.Request) {
	if e.handler == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	httpx.AddRequestId(r)
	resp, _ := e.handler(r)
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}
