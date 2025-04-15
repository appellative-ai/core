package rest

import (
	"context"
	"fmt"
	"net/http"
)

const (
	route1 = "route-1"
	route2 = "route-2"
)

func ExampleRouter() {
	r := NewRouter()

	r.Modify(route1, "https://www.google.com/search?q=golang", nil)
	route, ok := r.Lookup(route1)
	fmt.Printf("test: NewRouter(\"%v\") -> [name:%v] [uri:%v] [ex:%v] [ok:%v]\n", route1, route.Name, route.Uri, route.Ex, ok)

	r.Modify(route1, "https://search.yahoo.com/search?q=golang", nil)
	route, ok = r.Lookup(route1)
	fmt.Printf("test: NewRouter(\"%v\") -> [name:%v] [uri:%v] [ex:%v] [ok:%v]\n", route1, route.Name, route.Uri, route.Ex, ok)

	//Output:
	//test: NewRouter("route-1") -> [name:route-1] [uri:https://www.google.com/search?q=golang] [ex:<nil>] [ok:true]
	//test: NewRouter("route-1") -> [name:route-1] [uri:https://search.yahoo.com/search?q=golang] [ex:<nil>] [ok:true]

}

func RoutingLink(router *Router) func(next Exchange) Exchange {
	return func(_ Exchange) Exchange {
		return func(r *http.Request) (resp *http.Response, err error) {
			var (
				newRequest *http.Request
				route      *Route
				ok         bool
				name       = ""
			)
			// Logic to convert the incoming request into a route name and a new request
			if r.URL.Path == "/search/google" {
				name = route1
			} else {
				name = route2
			}
			if route, ok = router.Lookup(name); !ok {
				// Need to notify not found routing
				return &http.Response{StatusCode: http.StatusNotFound}, nil
			}
			switch name {
			case route1:
				newRequest = r.Clone(context.Background())
			case route2:
				newRequest = r.Clone(context.Background())

			}
			return route.Ex(newRequest)
		}
	}
}
