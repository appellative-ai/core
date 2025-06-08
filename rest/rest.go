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

// ExchangeLink - interface to link http Exchanges
type ExchangeLink func(next Exchange) Exchange

// Chainable - interface to link http Exchanges
type Chainable interface {
	Link(next Exchange) Exchange
}

// BuildChain - build a chain of http Exchanges - panic on nil or invalid type links
func BuildChain(links []any) Exchange {
	if len(links) == 0 {
		return nil
	}
	last := len(links) - 1
	// create the last link which is of type Exchange or Exchangeable
	head := exchangeLink(links[last])

	// build rest of chain
	for i := last - 1; i >= 0; i-- {
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
func exchangeLink(last any) Exchange {
	if ex, ok := last.(func(r *http.Request) (*http.Response, error)); ok {
		return ex
	}
	if exc, ok1 := last.(Exchangeable); ok1 {
		return exc.Exchange
	}
	panic(last)
}

// BuildMessagingChain - build a chain of messaging processing - panic on nil or invalid type links
func BuildMessagingChain(links []any) Exchange {
	if len(links) == 0 {
		return nil
	}
	last := len(links) - 1
	// create the last link which is of type Exchange or Exchangeable
	head := exchangeLink(links[last])

	// build chain
	for i := last; i >= 0; i-- {
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
