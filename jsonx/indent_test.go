package jsonx

import (
	"bytes"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"io"
)

const (
	customerAddr = "file://[cwd]/test/customer-address.txt"
)

func ExampleIndent() {
	buf, status := iox.ReadFile(customerAddr)
	fmt.Printf("test: iox.ReadFile() -> [status:%v] %v\n", status, string(buf))

	if status == nil {
		//fmt.Printf("test:")
		c := io.NopCloser(bytes.NewReader(buf))
		c2, status1 := Indent(c, nil, "", "  ")
		if status1 == nil {
			buf2, status2 := iox.ReadAll(c2, nil)
			fmt.Printf("test: Indent() -> [status:%v] %v\n", status2, string(buf2))
		}
	}

	//Output:
	//test: iox.ReadFile() -> [status:<nil>] [{"customer-id":"C001","created-ts":"0001-01-01T00:00:00Z","address-1":"1514 Cedar Ridge Road","address-2":"","city":"Vinton","state":"IA","postal-code":"52349","email":"before-email@hotmail.com"}]
	//test: Indent() -> [status:<nil>] [
	//  {
	//    "customer-id": "C001",
	//    "created-ts": "0001-01-01T00:00:00Z",
	//    "address-1": "1514 Cedar Ridge Road",
	//    "address-2": "",
	//    "city": "Vinton",
	//    "state": "IA",
	//    "postal-code": "52349",
	//    "email": "before-email@hotmail.com"
	//  }
	//]

}
