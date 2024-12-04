package jsonx

import (
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"io"
	"net/url"
	"os"
	"strings"
)

const (
	address1Url     = "file://[cwd]/test/address1.json"
	address2Url     = "file://[cwd]/test/address2.json"
	address2UrlGzip = "file://[cwd]/test/address2.gz"
	address3Url     = "file://[cwd]/test/address3.json"
	status504       = "file://[cwd]/test/status-504.json"
)

type newAddress struct {
	City    string
	State   string
	ZipCode string
}

// parseRaw - parse a raw Uri without error
func parseRaw(rawUri string) *url.URL {
	u, _ := url.Parse(rawUri)
	return u
}

func ExampleNew_String_Error() {
	_, status := New[newAddress]("", nil)
	fmt.Printf("test: New(\"\") -> [status:%v]\n", status)

	s := "https://www.google.com/search"
	_, status = New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	s = "file://[cwd]/test/address.txt"
	_, status = New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	//Output:
	//test: New("") -> [status:I/O Failure [open error: URL is empty: The system cannot find the file specified.]]
	//test: New(https://www.google.com/search) -> [status:I/O Failure [open error: scheme is invalid [https]: The system cannot find the file specified.]]
	//test: New(file://[cwd]/test/address.txt) -> [status:I/O Failure [open C:\Users\markb\GitHub\common\jsonx\test\address.txt: The system cannot find the file specified.]]

}

/*
func ExampleNew_String_Status() {
	s := StatusOKUri
	addr, status := New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	s = StatusNotFoundUri
	bytes, status0 := New[[]byte](s, nil)
	fmt.Printf("test: New(%v) -> [bytes:%v] [status:%v]\n", s, bytes, status0)

	s = StatusTimeoutUri
	addr, status = New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	// status from uri, generic type is ignored
	s = status504
	bytes1, status1 := New[[]byte](s, nil)
	fmt.Printf("test: New(%v) -> [bytes:%v] [status:%v]\n", s, bytes1, status1)

	//Output:
	//test: New(urn:status:ok) -> [addr:{  }] [status:OK]
	//test: New(urn:status:notfound) -> [bytes:[]] [status:Not Found]
	//test: New(urn:status:timeout) -> [addr:{  }] [status:Timeout]
	//test: New(file://[cwd]/test/status-504.jsonx) -> [bytes:[]] [status:Timeout [error 1]]

}


*/

func ExampleNew_String_URI() {
	// bytes
	s := address1Url
	//bytes, status := New[[]byte](s)
	//fmt.Printf("test: New(%v) -> [bytes:%v] [status:%v]\n", s, len(bytes), status)

	// type
	s = address1Url
	addr, status1 := New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/test/address1.json) -> [addr:{frisco texas 75034}] [status:OK]

}

func ExampleNew_URL_Error() {
	_, status := New[newAddress](nil, nil)
	fmt.Printf("test: New(\"\") -> [status:%v]\n", status)

	s := "https://www.google.com/search"
	_, status = New[newAddress](parseRaw(s), nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	s = "file://[cwd]/test/address.txt"
	_, status = New[newAddress](parseRaw(s), nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	//Output:
	//test: New("") -> [status:Invalid Argument [error: value parameter is nil]]
	//test: New(https://www.google.com/search) -> [status:I/O Failure [open error: scheme is invalid [https]: The system cannot find the file specified.]]
	//test: New(file://[cwd]/test/address.txt) -> [status:I/O Failure [open C:\Users\markb\GitHub\common\jsonx\test\address.txt: The system cannot find the file specified.]]

}

/*
func ExampleNew_URL_Status() {
	s := status504
	u, _ := url.Parse(s)

	addr, status0 := New[newAddress](u, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status0)

	//Output:
	//test: New(file://[cwd]/test/status-504.json) -> [addr:{  }] [status:Timeout [error 1]]

}


*/

func ExampleNew_URL() {
	s := address1Url
	u, _ := url.Parse(s)
	addr, status1 := New[newAddress](u, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/test/address1.json) -> [addr:{frisco texas 75034}] [status:OK]

}

func _ExampleZipFile() {
	status := iox.ZipFile(address2Url)

	fmt.Printf("test: ZipFile(\"\") -> [status:%v]\n", status)

	//Output:
	//test: ZipFile("") -> [status:OK]
}

func ExampleNew_Bytes() {
	s := address2Url
	buf, err := os.ReadFile(iox.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}
	addr, status := New[newAddress](buf, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	s = address2UrlGzip
	buf, err = os.ReadFile(iox.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}
	addr, status = New[newAddress](buf, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	//Output:
	//test: New(file://[cwd]/test/address2.json) -> [addr:{vinton iowa 52349}] [status:OK]
	//test: New(file://[cwd]/test/address2.gz) -> [addr:{vinton iowa 52349}] [status:OK]

}

func ExampleNew_Reader() {
	s := address2Url
	buf, err := os.ReadFile(iox.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}

	r := strings.NewReader(string(buf))
	addr, status1 := New[newAddress](r, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/test/address2.json) -> [addr:{vinton iowa 52349}] [status:OK]

}

func ExampleNew_ReadCloser() {
	s := address3Url
	buf, err := os.ReadFile(iox.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}

	body := io.NopCloser(strings.NewReader(string(buf)))
	addr, status1 := New[newAddress](body, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/test/address3.json) -> [addr:{forest city iowa 50436}] [status:OK]

}
