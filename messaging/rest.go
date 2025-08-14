package messaging

import (
	"fmt"
	"reflect"
)

// MessageLink - interface to link message handlers. Used in the collective repository
type MessageLink func(next *Message)

// Chainable - interface to create a link
type Chainable[T any] interface {
	Link(t T) T
}

// BuildMessagingChain - build messaging handler chain
func BuildMessagingChain(links []any) Handler {
	return BuildChain[Handler, Chainable[Handler]](links)
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
