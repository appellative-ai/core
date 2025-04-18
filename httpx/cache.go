package httpx

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"sync"
)

var (
	notFoundResponse = &http.Response{StatusCode: http.StatusNotFound}
)

type cacheT struct {
	body []byte
	resp *http.Response
}

type ResponseCache interface {
	Get(key string) *http.Response
	Put(key string, resp *http.Response) error
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
	if v2, ok1 := value.(cacheT); ok1 {
		v2.resp.Body = io.NopCloser(bytes.NewReader(v2.body))
		return v2.resp
	}
	return notFoundResponse
}

// Put - store response based on a URI, usually the URL
func (c *contentT) Put(uri string, resp *http.Response) error {
	if uri == "" || resp == nil {
		return errors.New("invalid argument: either uri is empty or http.Response is nil")
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := cacheT{
		body: buf,
		resp: resp,
	}
	c.m.Store(uri, data)
	return nil
}

// CreateResponse - create a response from a request
// TODO: may need to remove sensitive headers. See Go code for header cloning
func CreateResponse(r *http.Request) *http.Response {
	h := CloneHeader(r.Header)
	return &http.Response{StatusCode: http.StatusOK, Header: h, Body: r.Body}
}
