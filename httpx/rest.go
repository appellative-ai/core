package httpx

import (
	"net/http"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// Link - function to link http Exchanges
type Link func(next Exchange) Exchange

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
	if fn, ok := links[last].(func(next Exchange) Exchange); ok {
		head = fn(nil)
	} else {
		if c, ok1 := links[last].(Chainable); ok1 {
			head = c.Link(nil)
		} else {
			panic(links[last])
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
