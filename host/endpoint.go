package host

import (
	"github.com/behavioral-ai/core/httpx"
	"net/http"
)

// ExchangeHandler - http exchange handler interface
type ExchangeHandler interface {
	Exchange(w http.ResponseWriter, r *http.Request)
}

type endpoint struct {
	handler httpx.Exchange
}

func NewEndpoint(handler httpx.Exchange) ExchangeHandler {
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
