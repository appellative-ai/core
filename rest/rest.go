package rest

import (
	"net/http"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// ExchangeHandler - extend the http.HandlerFunc to include the http.Response
type ExchangeHandler func(w http.ResponseWriter, req *http.Request, resp *http.Response)

// Exchangeable - interface to http Exchanges
type Exchangeable interface {
	Exchange(r *http.Request) (*http.Response, error)
}

// Chainable - interface to link http Exchanges
type Chainable interface {
	Link(next Exchange) Exchange
}

// BuildChain - build a chain of http Exchanges - panic on nil or invalid type links
func BuildChain(links ...any) Exchange {
	if len(links) == 0 {
		return nil
	}

	// create the last link which is of type Exchange or Exchangeable
	head := exchangeLink(links)

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

// exchangeLink - last link needs to be of type Exchange or Exchangeable
func exchangeLink(links []any) Exchange {
	last := len(links) - 1
	if ex, ok := links[last].(func(r *http.Request) (*http.Response, error)); ok {
		return ex
	}
	if exc, ok1 := links[last].(Exchangeable); ok1 {
		return exc.Exchange
	}
	/*
		if fn, ok2 := links[last].(func(next Exchange) Exchange); ok2 {
			return fn(nil)
		}
		if c, ok3 := links[last].(Chainable); ok3 {
			return c.Link(nil)
		}

	*/
	panic(links[last])
}
