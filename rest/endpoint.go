package rest

import "net/http"

type Endpoint struct {
	Handler ExchangeHandler
	Chain   Exchange
	Init    func(r *http.Request)
}

func NewEndpoint(handler ExchangeHandler, init func(r *http.Request), chain Exchange) *Endpoint {
	e := new(Endpoint)
	e.Handler = handler
	e.Chain = chain
	e.Init = init
	return e
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e.Handler == nil || e.Chain == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if e.Init != nil {
		e.Init(r)
	}
	resp, _ := e.Chain(r)
	e.Handler(w, r, resp)
}
