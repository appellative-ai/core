package io

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

//go:embed iotest
var tf embed.FS

const (
	testResponseTxt  = "file://[cwd]/iotest/test-response.txt"
	testResponse2Txt = "file://[cwd]/iotest/test-response2.txt"

	helloWorldTxt  = "file://[cwd]/iotest/hello-world.txt"
	helloWorldGzip = "file://[cwd]/iotest/hello-world.gz"

	testResponseGzip = "file://[cwd]/iotest/test-response.gz"

	address1Url = "file://[cwd]/iotest/address1.json"
	address2Url = "file://[cwd]/iotest/address2.json"
	address3Url = "file://[cwd]/iotest/address3.json"
	status504   = "file://[cwd]/iotest/status-504.json"
)

// parseRaw - parse a raw Uri without error
func parseRaw(rawUri string) *url.URL {
	u, _ := url.Parse(rawUri)
	return u
}

func ExampleDirFS() {
	dir := "file:///c:/Users/markb/GitHub/core/io/iotest"
	fileSystem := DirFS(dir)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path != "." {

			buf, err1 := fs.ReadFile(fileSystem, path)
			fmt.Printf("test: fs.ReadFile() -> [err:%v] [%v] %v\n", err1, path, buf != nil)
		}
		return nil
	})

	//Output:
	//test: fs.ReadFile() -> [err:<nil>] [address1.json] true
	//test: fs.ReadFile() -> [err:<nil>] [address2.json] true
	//test: fs.ReadFile() -> [err:<nil>] [address3.json] true
	//test: fs.ReadFile() -> [err:<nil>] [hello-world.gz] true
	//test: fs.ReadFile() -> [err:<nil>] [hello-world.txt] true
	//test: fs.ReadFile() -> [err:<nil>] [status-504.json] true
	//test: fs.ReadFile() -> [err:<nil>] [test-response.gz] true
	//test: fs.ReadFile() -> [err:<nil>] [test-response.txt] true
	//test: fs.ReadFile() -> [err:<nil>] [test-response2.gz] true
	//test: fs.ReadFile() -> [err:<nil>] [test-response2.txt] true

}

func ExampleDirFS_Failure() {
	dir := "file:///c:/Users/markb/GitHub/core/io/invalid"
	fileSystem := DirFS(dir)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("test: DirFS_Failure() -> [err:%v]\n", err) //log.Fatal(err)
		}
		if path != "." {

			buf, err1 := fs.ReadFile(fileSystem, path)
			fmt.Printf("test: fs.ReadFile() -> [err:%v] [%v] %v\n", err1, path, buf != nil)
		}
		return nil
	})

	//Output:
	//test: DirFS_Failure() -> [err:CreateFile .: The system cannot find the file specified.]

}

func Example_FileNameError() {
	//s := "file://[cwd]/test/test-response.txt"
	//u, err := url.Parse(s)
	//fmt.Printf("test: url.Parse(%v) -> [err:%v]\n", s, err)

	var t any
	name := FileName(t)
	fmt.Printf("test: FileName(nil) -> [type:%v] [url:%v]\n", reflect.TypeOf(t), name)

	s := ""
	name = FileName(s)
	fmt.Printf("test: FileName(\"\") -> [type:%v] [url:%v]\n", reflect.TypeOf(s), name)

	s = "https://www.google.com/search?q=golang"
	name = FileName(s)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(s), name)

	s = "https://www.google.com/search?q=golang"
	u := parseRaw(s)
	name = FileName(u)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(u), name)

	req, _ := http.NewRequest("", s, nil)
	name = FileName(req)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(req), name)

	s = "file://[cwd]/iotest/test-response.txt"
	req, _ = http.NewRequest("", s, nil)
	name = FileName(req)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(req), name)

	//Output:
	//test: FileName(nil) -> [type:<nil>] [url:error: URL is nil]
	//test: FileName("") -> [type:string] [url:error: URL is empty]
	//test: FileName(https://www.google.com/search?q=golang) -> [type:string] [url:error: scheme is invalid [https]]
	//test: FileName(https://www.google.com/search?q=golang) -> [type:*url.URL] [url:error: scheme is invalid [https]]
	//test: FileName(https://www.google.com/search?q=golang) -> [type:*http.Request] [url:error: invalid URL type: *http.Request]
	//test: FileName(file://[cwd]/iotest/test-response.txt) -> [type:*http.Request] [url:error: invalid URL type: *http.Request]

}

func Example_FileName() {
	s := "file://[cwd]/iotest/test-response.txt"
	u, err := url.Parse(s)
	fmt.Printf("test: url.Parse(%v) -> [err:%v]\n", s, err)

	name := FileName(s)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(s), name)

	name = FileName(u)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(u), name)

	s = "file:///c:/Users/markb/GitHub/core/io/iotest/test-response.txt"
	name = FileName(s)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(s), name)

	u, err = url.Parse(s)
	name = FileName(u)
	fmt.Printf("test: FileName(%v) -> [type:%v] [url:%v]\n", s, reflect.TypeOf(u), name)

	//Output:
	//test: url.Parse(file://[cwd]/iotest/test-response.txt) -> [err:<nil>]
	//test: FileName(file://[cwd]/iotest/test-response.txt) -> [type:string] [url:C:\Users\markb\GitHub\core\io\iotest\test-response.txt]
	//test: FileName(file://[cwd]/iotest/test-response.txt) -> [type:*url.URL] [url:C:\Users\markb\GitHub\core\io\iotest\test-response.txt]
	//test: FileName(file:///c:/Users/markb/GitHub/core/io/iotest/test-response.txt) -> [type:string] [url:c:\Users\markb\GitHub\core\io\iotest\test-response.txt]
	//test: FileName(file:///c:/Users/markb/GitHub/core/io/iotest/test-response.txt) -> [type:*url.URL] [url:c:\Users\markb\GitHub\core\io\iotest\test-response.txt]

}

func Example_OSReadFile() {
	s := "file://[cwd]/iotest/test-response.txt"
	u, _ := url.Parse(s)
	buf, err := os.ReadFile(FileName(u))
	fmt.Printf("test: os.ReadFile(%v) -> [err:%v] [buf:%v]\n", s, err, len(buf))

	s = "file:///c:/Users/markb/GitHub/core/io/iotest/test-response.txt"
	u, _ = url.Parse(s)
	buf, err = os.ReadFile(FileName(u))
	fmt.Printf("test: os.ReadFile(%v) -> [err:%v] [buf:%v]\n", s, err, len(buf))

	//Output:
	//test: os.ReadFile(file://[cwd]/iotest/test-response.txt) -> [err:<nil>] [buf:188]
	//test: os.ReadFile(file:///c:/Users/markb/GitHub/core/io/iotest/test-response.txt) -> [err:<nil>] [buf:188]

}

func ExampleReadFile() {
	s := status504
	buf, status := ReadFile(s)
	fmt.Printf("test: ReadFile(%v) -> [type:%v] [buf:%v] [status:%v]\n", s, reflect.TypeOf(s), len(buf), status)

	s = address1Url
	buf, status = ReadFile(s)
	fmt.Printf("test: ReadFile(%v) -> [type:%v] [buf:%v] [status:%v]\n", s, reflect.TypeOf(s), len(buf), status)

	s = status504
	u := parseRaw(s)
	buf, status = ReadFile(u.String())
	fmt.Printf("test: ReadFile(%v) -> [type:%v] [buf:%v] [status:%v]\n", s, reflect.TypeOf(u), len(buf), status)

	s = address1Url
	u = parseRaw(s)
	buf, status = ReadFile(u.String())
	fmt.Printf("test: ReadFile(%v) -> [type:%v] [buf:%v] [status:%v]\n", s, reflect.TypeOf(u), len(buf), status)

	//Output:
	//test: ReadFile(file://[cwd]/iotest/status-504.json) -> [type:string] [buf:82] [status:<nil>]
	//test: ReadFile(file://[cwd]/iotest/address1.json) -> [type:string] [buf:68] [status:<nil>]
	//test: ReadFile(file://[cwd]/iotest/status-504.json) -> [type:*url.URL] [buf:82] [status:<nil>]
	//test: ReadFile(file://[cwd]/iotest/address1.json) -> [type:*url.URL] [buf:68] [status:<nil>]

}

func ExampleReadFileWithEncoding() {
	buf, status := ReadFileWithEncoding(helloWorldGzip, nil)
	fmt.Printf("test: ReadFileWithEncoding(\"%v\",nil) -> [buf:%v] [status:%v]\n", helloWorldGzip, string(buf), status)

	h := make(http.Header)
	h.Set(ContentEncoding, GzipEncoding)
	buf, status = ReadFileWithEncoding(helloWorldGzip, h)
	fmt.Printf("test: ReadFileWithEncoding(\"%v\",h) -> [buf:%v] [status:%v]\n", helloWorldGzip, string(buf), status)

	buf, status = ReadFileWithEncoding(helloWorldTxt, nil)
	fmt.Printf("test: ReadFileWithEncoding(\"%v\",nil) -> [buf:%v] [status:%v]\n", helloWorldTxt, string(buf), status)

	//Output:
	//test: ReadFileWithEncoding("file://[cwd]/iotest/hello-world.gz",nil) -> [buf:Hello World!!] [status:<nil>]
	//test: ReadFileWithEncoding("file://[cwd]/iotest/hello-world.gz",h) -> [buf:Hello World!!] [status:<nil>]
	//test: ReadFileWithEncoding("file://[cwd]/iotest/hello-world.txt",nil) -> [buf:Hello World!!] [status:<nil>]

}

func ExampleReadFileEmbedded() {
	name := "file:///f:/iotest/hello-world.txt"

	bytes, status := ReadFile(name)
	fmt.Printf("test: ReadFileEmbedded(\"%v\") -> [buf:%v] [status:%v]\n", name, string(bytes), status)

	Mount(tf)
	Mount(tf)

	name2 := "file:///f:/iotest/invalid-file-name"
	bytes, status = ReadFile(name2)
	fmt.Printf("test: ReadFileEmbedded(\"%v\") -> [buf:%v] [status:%v]\n", name2, string(bytes), status)

	bytes, status = ReadFile(name)
	fmt.Printf("test: ReadFileEmbedded(\"%v\") -> [buf:%v] [status:%v]\n", name, string(bytes), status)

	//Output:
	//test: ReadFileEmbedded("file:///f:/iotest/hello-world.txt") -> [buf:] [status:open iotest/hello-world.txt: file does not exist]
	//error: file system is already mounted
	//test: ReadFileEmbedded("file:///f:/iotest/invalid-file-name") -> [buf:] [status:open iotest/invalid-file-name: file does not exist]
	//test: ReadFileEmbedded("file:///f:/iotest/hello-world.txt") -> [buf:Hello World!!] [status:<nil>]

}
