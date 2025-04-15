package rest

import (
	"fmt"
	"net/http"
)

func ExampleNewRouteMessage() {
	m := NewRouteMessage("test-route", "https://www.google.com/search?q=golang", func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusTeapot}, nil
	})
	r, ok := RouteContent(m)
	fmt.Printf("test: NewRouteMessage() [name:%v] [uri:%v] [ok:%v]\n", r.Name, r.Uri, ok)

	//Output:
	//test: NewRouteMessage() [name:test-route] [uri:https://www.google.com/search?q=golang] [ok:true]

}
