package rest

import "net/http"

type Endpoint interface {
	Pattern() string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
type endpoint struct {
	pattern string
	handler ExchangeHandler
	chain   Exchange
	init    func(r *http.Request)
}

func NewEndpoint(pattern string, handler ExchangeHandler, init func(r *http.Request), chain Exchange) Endpoint {
	e := new(endpoint)
	e.pattern = pattern
	e.handler = handler
	e.chain = chain
	e.init = init
	return e
}

func (e *endpoint) Pattern() string {
	return e.pattern
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e.handler == nil || e.chain == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if e.init != nil {
		e.init(r)
	}
	resp, _ := e.chain(r)
	e.handler(w, r, resp)
}
