package httpx

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func ExampleEmptyReader() {
	r := emptyReader

	buf, err := io.ReadAll(r)
	fmt.Printf("test: emptyReader() -> [buf:%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: emptyReader() -> [buf:] [err:<nil>]

}

func ExampleReadAll() {
	s := "this is new content"
	r1 := bytes.NewReader([]byte(s))
	r2 := bytes.NewReader([]byte(s))

	go func() {
		buf, err := io.ReadAll(r1)
		fmt.Printf("test: io.ReadAll(r1) -> [buf:%v] [err:%v]\n", string(buf), err)
	}()

	buf, err := io.ReadAll(r2)
	fmt.Printf("test: io.ReadAll(r2) -> [buf:%v] [err:%v]\n", string(buf), err)

	time.Sleep(time.Second * 2)
	//Output:
	//test: io.ReadAll(r2) -> [buf:this is new content] [err:<nil>]
	//test: io.ReadAll(r1) -> [buf:this is new content] [err:<nil>]

}
