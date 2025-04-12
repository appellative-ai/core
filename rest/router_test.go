package rest

import (
	"context"
	"net/http"
)

const (
	route1 = "route-1"
	route2 = "route-2"
)

func ExampleRouter() {

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
