package rest

import (
	"errors"
)

type Route struct {
	Name string
	Uri  string
	Ex   Exchange
}

func NewRoute(name, uri string, ex Exchange) *Route {
	r := new(Route)
	r.Name = name
	r.Uri = uri
	r.Ex = ex
	return r
}

type Router struct {
	table *routeTable
}

func NewRouter() *Router {
	r := new(Router)
	r.table = newRouteTable()
	//for _, r := range routes {
	//	router.table.put(r)
	//}
	return r
}

func (r *Router) Modify(name, uri string, ex Exchange) error {
	if name == "" {
		return errors.New("route name is empty")
	}
	rt, ok := r.table.get(name)
	if ok {
		if uri != "" {
			rt.Uri = uri
		}
		if ex != nil {
			rt.Ex = ex
		}
		r.table.put(rt)
	} else {
		r.table.put(NewRoute(name, uri, ex))
	}
	return nil
}

func (r *Router) Lookup(name string) (*Route, bool) {
	if route, ok := r.table.get(name); ok {
		return route, true
	}
	return nil, false
}
