package rest

type Route struct {
	Name string
	Ex   Exchange
}

func NewRoute(name string, ex Exchange) *Route {
	r := new(Route)
	r.Name = name
	r.Ex = ex
	return r
}

type Router struct {
	table routeTable
}

func NewRouter(routes ...*Route) *Router {
	router := new(Router)
	for _, r := range routes {
		router.table.put(r)
	}
	return router
}

func (r *Router) Lookup(name string) (*Route, bool) {
	if route, ok := r.table.get(name); ok {
		return route, true
	}
	return nil, false
}
