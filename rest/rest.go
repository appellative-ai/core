package rest

import (
	"net/http"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// ExchangeHandler - extend the http.HandlerFunc to include the http.Response
type ExchangeHandler func(w http.ResponseWriter, req *http.Request, resp *http.Response)

// Chainable - interface to link http Exchanges
type Chainable interface {
	Link(next Exchange) Exchange
}

// BuildChain - build a chain of http Exchanges - panic on nil or invalid type links
func BuildChain(links ...any) Exchange {
	if len(links) == 0 {
		return nil
	}
	var head Exchange

	// initialize head to last item
	last := len(links) - 1
	// Allow last item to be an Exchange and not linkable
	if ex, ok := links[last].(func(r *http.Request) (*http.Response, error)); ok {
		head = ex
	} else {
		if fn, ok1 := links[last].(func(next Exchange) Exchange); ok1 {
			head = fn(nil)
		} else {
			if c, ok2 := links[last].(Chainable); ok2 {
				head = c.Link(nil)
			} else {
				panic(links[last])
			}
		}
	}

	// build rest of chain
	for i := len(links) - 2; i >= 0; i-- {
		if fn, ok := links[i].(func(next Exchange) Exchange); ok {
			head = fn(head)
			continue
		}
		if c, ok := links[i].(Chainable); ok {
			head = c.Link(head)
			continue
		}
		panic(links[i])
	}
	return head
}
