package rest

import (
	"sync"
)

type routeTable struct {
	m *sync.Map
}

func newRouteTable() *routeTable {
	t := new(routeTable)
	t.m = new(sync.Map)
	return t
}

func (t *routeTable) get(name string) (*Route, bool) {
	value, ok := t.m.Load(name)
	if !ok {
		return nil, false
	}
	if value1, ok1 := value.(*Route); ok1 {
		return value1, true
	}
	return nil, false
}

func (t *routeTable) put(route *Route) {
	if route != nil && route.Name != "" {
		t.m.Store(route.Name, route)
	}
}
