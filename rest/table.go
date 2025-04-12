package rest

import (
	"sync"
)

type routeTable struct {
	m *sync.Map
}

func newRouteTable() *routeTable {
	r := new(routeTable)
	r.m = new(sync.Map)
	return r
}

func (c *routeTable) get(name string) (*Route, bool) {
	value, ok := c.m.Load(name)
	if !ok {
		return nil, false
	}
	if value1, ok1 := value.(*Route); ok1 {
		return value1, true
	}
	return nil, false
}

func (c *routeTable) put(route *Route) {
	if route != nil && route.Name != "" {
		c.m.Store(route.Name, route)
	}
}
