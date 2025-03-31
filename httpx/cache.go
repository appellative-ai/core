package httpx

import (
	"net/http"
	"sync"
)

var (
	notFoundResponse = &http.Response{StatusCode: http.StatusNotFound}
)

type ResponseCache interface {
	Get(key string) *http.Response
	Put(key string, resp *http.Response)
}

type contentT struct {
	m *sync.Map
}

func NewResponseCache() ResponseCache {
	c := new(contentT)
	c.m = new(sync.Map)
	return newCache()
}

func newCache() *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	return c
}

// Get - load a response based on a URI, usually the URL
func (c *contentT) Get(uri string) *http.Response {
	value, ok := c.m.Load(uri)
	if !ok {
		return notFoundResponse
	}
	if r, ok1 := value.(*http.Response); ok1 {
		return r
	}
	return notFoundResponse
}

// Put - store response based on a URI, usually the URL
func (c *contentT) Put(uri string, resp *http.Response) {
	c.m.Store(uri, resp)
}

// CreateResponse - create a response from a request
// TODO: may need to remove sensitive headers. See Go code for header cloning
func CreateResponse(r *http.Request) *http.Response {
	h := CloneHeader(r.Header)
	return &http.Response{StatusCode: http.StatusOK, Header: h, Body: r.Body}
}
