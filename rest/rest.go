package rest

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// ExchangeHandler - extend the http.HandlerFunc to include the http.Response
type ExchangeHandler func(w http.ResponseWriter, req *http.Request, resp *http.Response)

// ExchangeLink - interface to link http Exchanges. Used in the collective repository
type ExchangeLink func(next Exchange) Exchange

// Chainable - interface to create a link
type Chainable[T any] interface {
	Link(t T) T
}

// BuildExchangeChain - build Exchange chain
func BuildExchangeChain(links []any) Exchange {
	return BuildChain[Exchange, Chainable[Exchange]](links)
}

// BuildReceiverChain - build messaging Receiver chain
func BuildReceiverChain(links []any) messaging.Receiver {
	return BuildChain[messaging.Receiver, Chainable[messaging.Receiver]](links)
}

// BuildChain - build a chain of links - panic on nil or invalid type links
func BuildChain[T any, U Chainable[T]](links []any) (head T) {
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
		panic(links[i])
	}
	return head
}

/*
// Exchangeable - interface to http Exchanges
type Exchangeable interface {
	Exchange(r *http.Request) (*http.Response, error)
}
*/
