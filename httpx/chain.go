package httpx

import (
	"net/http"
)

// Agent - adds a chainable exchange method
/*
type Agent interface {
	messaging.Agent
	Link(next Exchange) Exchange
}

*/

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

type Link func(next Exchange) Exchange

// Chainable - provides a link in a chain of http Exchanges
type Chainable interface {
	Link(next Exchange) Exchange
}

func BuildChainT(ex ...Chainable) Exchange {
	if len(ex) == 0 {
		return nil
	}
	var head Exchange

	for i := len(ex) - 1; i >= 0; i-- {
		if i == len(ex)-1 {
			head = ex[i].Link(nil)
		} else {
			head = ex[i].Link(head)
		}
	}
	return head
}

func BuildChain(ex ...Link) Exchange {
	if len(ex) == 0 {
		return nil
	}
	var head Exchange

	for i := len(ex) - 1; i >= 0; i-- {
		if i == len(ex)-1 {
			head = ex[i](nil)
		} else {
			head = ex[i](head)
		}
	}
	return head
}
