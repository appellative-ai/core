package httpx

import (
	"fmt"
	"io"
)

func ExampleEmptyReader() {
	r := emptyReader

	buf, err := io.ReadAll(r)
	fmt.Printf("test: emptyReader() -> [buf:%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: emptyReader() -> [buf:] [err:<nil>]

}
