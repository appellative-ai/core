package rest

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"reflect"
)

// Micro-REST

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// ExchangeLink - interface to link http Exchanges. Used in the collective repository
type ExchangeLink func(next Exchange) Exchange

// MessageLink - interface to link message handlers. Used in the collective repository
type MessageLink func(next *messaging.Message)

// Chainable - interface to create a link
type Chainable[T any] interface {
	Link(t T) T
}

// BuildExchangeChain - build Exchange chain
func BuildExchangeChain(links []any) Exchange {
	return BuildChain[Exchange, Chainable[Exchange]](links)
}

// BuildMessagingChain - build messaging handler chain
func BuildMessagingChain(links []any) messaging.Handler {
	return BuildChain[messaging.Handler, Chainable[messaging.Handler]](links)
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
		panic(fmt.Sprintf("invalid link type: %v", reflect.TypeOf(links[i])))
	}
	return head
}

/*
// Exchangeable - interface to http Exchanges
type Exchangeable interface {
	Exchange(r *http.Request) (*http.Response, error)
}
*/
