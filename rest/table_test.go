package rest

import (
	"fmt"
	"net/http"
)

func ExampleRouteTable() {
	rt1 := "route-1"
	rt2 := "route-2"
	t := newRouteTable()

	t.put(NewRoute(rt1, "https://www.google.com/search?q=golang", nil))
	route, ok := t.get(rt1)
	fmt.Printf("test: newRouteTable() -> [name:%v] [uri:%v] [ex:%v] [ok:%v]\n", route.Name, route.Uri, route.Ex, ok)

	t.put(NewRoute(rt2, "https://search.yahoo.com/search?q=golang", nil))
	route, ok = t.get(rt2)
	fmt.Printf("test: newRouteTable() -> [name:%v] [uri:%v] [ex:%v] [ok:%v]\n", route.Name, route.Uri, route.Ex, ok)

	t.put(NewRoute(rt1, "https://duckduckgo.com/search?q=golang", func(r *http.Request) (*http.Response, error) {
		return nil, nil
	}))
	route, ok = t.get(rt1)
	fmt.Printf("test: newRouteTable() -> [name:%v] [uri:%v] [ex:%v] [ok:%v]\n", route.Name, route.Uri, route.Ex != nil, ok)

	//Output:
	//test: newRouteTable() -> [name:route-1] [uri:https://www.google.com/search?q=golang] [ex:<nil>] [ok:true]
	//test: newRouteTable() -> [name:route-2] [uri:https://search.yahoo.com/search?q=golang] [ex:<nil>] [ok:true]
	//test: newRouteTable() -> [name:route-1] [uri:https://duckduckgo.com/search?q=golang] [ex:true] [ok:true]

}
