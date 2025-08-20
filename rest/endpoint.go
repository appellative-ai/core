package rest

import "net/http"

// HttpHandler - extend the http.HandlerFunc to include the http.Response
type HttpHandler func(w http.ResponseWriter, req *http.Request, resp *http.Response)

type Endpoint interface {
	Pattern() string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type endpoint struct {
	pattern string
	handler HttpHandler
	chain   Exchange
	init    func(r *http.Request)
}

func NewEndpoint(pattern string, handler HttpHandler, init func(r *http.Request), operatives []any) Endpoint {
	e := new(endpoint)
	e.pattern = pattern
	e.handler = handler
	e.chain = BuildNetwork(operatives)
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
