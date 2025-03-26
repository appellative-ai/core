package httpx

import "net/http"

// Exchange - http exchange
type Exchange func(r *http.Request) (*http.Response, error)

// Chainable - provides a link in a chain of Exchanges
type Chainable interface {
	Link(next Exchange) Exchange
}

func BuildChain(ex ...Chainable) Exchange {
	if len(ex) == 0 {
		return nil
	}
	var head Exchange

	for i := len(ex) - 1; i >= 0; i-- {
		f := ex[i]
		if i == len(ex)-1 {
			head = f.Link(nil)
		} else {
			head = f.Link(head)
		}
	}
	return head
}
