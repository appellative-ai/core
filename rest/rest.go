package rest

import (
	"net/http"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// ExchangeHandler - extend the http.HandlerFunc to include the http.Response
type ExchangeHandler func(w http.ResponseWriter, req *http.Request, resp *http.Response)

/*
// Exchangeable - interface to http Exchanges
type Exchangeable interface {
	Exchange(r *http.Request) (*http.Response, error)
}

// ExchangeLink - interface to link http Exchanges
type ExchangeLink func(next Exchange) Exchange
*/

// Chainable - interface to create a link
type Chainable[T any] interface {
	Link(t T) T
}

// BuildChain - build a chain of links - panic on nil or invalid type links
func BuildChain[T any, U Chainable[T]](links []any) (head T) {
	if len(links) == 0 {
		return head
	}
	for i := len(links) - 1; i >= 0; i-- {
		// Check for a next function
		if fn, ok := links[i].(func(next T) T); ok {
			head = fn(head)
			continue
		}
		// Check for a Chainable interface
		if c, ok := links[i].(U); ok {
			head = c.Link(head)
			continue
		}
		panic(links[i])
	}
	return head
}
