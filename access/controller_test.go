package access

import (
	"fmt"
	"reflect"
	"time"
)

func ExampleController() {
	var timeout time.Duration
	var limit float64
	var burst int
	var redirect int

	timeout = time.Millisecond * 2000
	limit = 100
	burst = 10
	redirect = 40

	metrics := []any{timeout, limit, burst, redirect}

	for _, v := range metrics {
		i := reflect.TypeOf(v)
		fmt.Printf("test: Metrics() -> [name:%v]\n", i.Name())
		fmt.Printf("test: Metrics() -> [type:%v] [value:%v]\n", reflect.TypeOf(v), v)
	}

	//Output:
	//fail

}
