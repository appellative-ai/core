package rest

import "net/http"

type endpoint struct {
	handler ExchangeHandler
	chain   Exchange
	init    func(r *http.Request)
}

func NewEndpoint(handler ExchangeHandler, init func(r *http.Request), chain Exchange) http.Handler {
	e := new(endpoint)
	e.handler = handler
	e.chain = chain
	if init == nil {
		e.init = func(r *http.Request) {}
	} else {
		e.init = init
	}
	return e
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e.handler == nil || e.chain == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e.init(r)
	resp, _ := e.chain(r)
	e.handler(w, r, resp)
}
