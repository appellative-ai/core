package host

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleNewMap() {
	m := NewSyncMap[string, messaging.NewAgent]()
	name := ""
	t := m.Load("")
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, func() messaging.Agent { return nil })
	t = m.Load(name)
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t != nil)

	//Output:
	//test:  get("") -> <nil>
	//test:  get("common:core:ctor/test") -> true

}
