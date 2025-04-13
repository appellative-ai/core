package host

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"net/http"
)

/*
// ExchangeHandler2 - http exchange handler interface
type ExchangeHandler2 interface {
	Exchange(w http.ResponseWriter, r *http.Request)
}

type endpoint2 struct {
	handler rest.Exchange
}

func NewEndpoint2(handler rest.Exchange) ExchangeHandler2 {
	e := new(endpoint2)
	e.handler = handler
	return e
}

func (e *endpoint2) Exchange(w http.ResponseWriter, r *http.Request) {
	if e.handler == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	httpx.AddRequestId(r)
	resp, _ := e.handler(r)
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}


*/

func ExchangeHandler(w http.ResponseWriter, req *http.Request, resp *http.Response) {
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, req.Header)
}

func Init(r *http.Request) {
	httpx.AddRequestId(r)
}

func NewEndpoint(links ...any) *rest.Endpoint {
	chain := rest.BuildChain(AccessLogLink, AuthorizationLink, links)
	return rest.NewEndpoint(ExchangeHandler, Init, chain)
}
