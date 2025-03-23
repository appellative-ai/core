package httpx

import (
	"fmt"
	"net/http"
)

func ExampleCopy() {
	h := make(http.Header)
	h.Add("key-1", "value-1")
	h.Add("key-2", "value-2")
	h.Add("key-3", "value-3")

	h2 := Copy(h)

	fmt.Printf("test: Copy() -> %v\n", h2)

	//Output:
	//fail

}
