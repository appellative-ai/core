package rest

import (
	"fmt"
	"net/http"
	"reflect"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// ExchangeLink - interface to link http Exchanges. Used in the collective exchange
type ExchangeLink func(next Exchange) Exchange

// Chainable - interface to create a link
type Chainable[T any] interface {
	Link(t T) T
}

// BuildExchangeChain - build Exchange chain
func BuildExchangeChain(links []any) Exchange {
	return BuildNetwork[Exchange, Chainable[Exchange]](links)
}

// BuildNetwork - build a chain of links - panic on nil or invalid type links
func BuildNetwork[T any, U Chainable[T]](links []any) (head T) {
	if len(links) == 0 {
		panic("error: chain links slice is nil")
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
		panic(fmt.Sprintf("invalid link type: %v", reflect.TypeOf(links[i])))
	}
	return head
}
