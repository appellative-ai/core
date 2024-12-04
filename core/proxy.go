package core

import (
	"errors"
	"fmt"
	uri2 "github.com/behavioral-ai/core/uri"
	"net/http"
	"sync"
)

// ExchangeProxy - key value pairs of a domain -> HttpExchange
type ExchangeProxy struct {
	m *sync.Map
}

// NewExchangeProxy - create a new Exchange Proxy
func NewExchangeProxy() *ExchangeProxy {
	p := new(ExchangeProxy)
	p.m = new(sync.Map)
	return p
}

func (p *ExchangeProxy) Register(domain string, handler HttpExchange) error {
	if len(domain) == 0 {
		return errors.New("invalid argument: domain is empty")
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("invalid argument: HTTP Exchange is nil for domain : [%v]", domain))
	}
	_, ok1 := p.m.Load(domain)
	if ok1 {
		return errors.New(fmt.Sprintf("invalid argument: HTTP Exchange already exists for domain : [%v]", domain))
	}
	p.m.Store(domain, handler)
	return nil
}

// LookupByRequest - find an HttpExchange from a request
func (p *ExchangeProxy) LookupByRequest(req *http.Request) HttpExchange {
	if req == nil || req.URL == nil {
		return nil
	}
	// Try host first
	ex := p.Lookup(req.Host)
	if ex != nil {
		return ex
	}

	// Default to embedded domain
	parsed := uri2.Uproot(req.URL.Path)
	if parsed.Valid {
		ex = p.Lookup(parsed.Domain)
	}
	return ex
}

// Lookup - get an HttpExchange from the proxy, using an domain as a key
func (p *ExchangeProxy) Lookup(domain string) HttpExchange {
	v, ok := p.m.Load(domain)
	if !ok {
		return nil //, errors.New(fmt.Sprintf("error: proxyLookupBydomain() HTTP handler does not exist: [%v]", domain))
	}
	if handler, ok1 := v.(HttpExchange); ok1 {
		return handler
	}
	return nil
}
