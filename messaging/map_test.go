package messaging

import (
	"fmt"
)

func ExampleNewMap() {
	m := NewSyncMap[string, NewAgentFunc]()
	name := ""
	t := m.Load("")
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, func() Agent { return nil })
	t = m.Load(name)
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t != nil)

	//Output:
	//test:  get("") -> <nil>
	//test:  get("common:core:ctor/test") -> true

}
