package rest

import (
	"fmt"
	"net/http"
)

func ExampleNewEndpoint() {
	e := NewEndpoint(nil, nil, nil)

	_, ok := any(e).(http.Handler)
	fmt.Printf("test: NewEndpoint() -> [%v] [ServeHTTP:%v]\n", e, ok)

	//Output:
	//test: NewEndpoint() -> [&{<nil> <nil> <nil>}] [ServeHTTP:true]

}
