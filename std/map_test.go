package std

import (
	"fmt"
)

type agent interface {
	Name() string
}

type newAgentFunc func() agent

func ExampleNewMap() {
	m := NewSyncMap[string, newAgentFunc]()
	name1 := ""
	t := m.Load("")
	fmt.Printf("test:  get(\"%v\") -> %v\n", name1, t)

	name1 = "common:core:ctor/test"
	m.Store(name1, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name1, func() agent { return nil })
	t = m.Load(name1)
	fmt.Printf("test:  get(\"%v\") -> %v\n", name1, t != nil)

	//Output:
	//test:  get("") -> <nil>
	//test:  get("common:core:ctor/test") -> true

}
