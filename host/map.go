package host

import (
	"sync"
)

// MapT - map type
type MapT[T, U any] struct {
	m *sync.Map
}

// NewSyncMap - create a new map
func NewSyncMap[T, U any]() *MapT[T, U] {
	c := new(MapT[T, U])
	c.m = new(sync.Map)
	return c
}

func (m *MapT[T, U]) Load(t T) (u U) {
	v, ok := m.m.Load(t)
	if !ok {
		return u
	}
	if v1, ok1 := v.(U); ok1 {
		return v1
	}
	return u
}

func (m *MapT[T, U]) Store(t T, u U) {
	m.m.Store(t, u)
}
