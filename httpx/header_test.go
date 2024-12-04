package httpx

import (
	"fmt"
	"net/http/httptest"
)

/*
func ExampleSelect() {
	h := http.Header{}

	resp := http.Response{Header: http.Header{}}
	resp.Header.Add("key", "value")
	resp.Header.Add("key1", "value1")
	resp.Header.Add("key2", "value2")

	CreateHeaders_OLD(h, &resp, "key", "key2")
	fmt.Printf("test: CreateHeaders() -> %v\n", h)

	h = http.Header{}
	CreateHeaders_OLD(h, &resp, "*")
	fmt.Printf("test: CreateHeaders() -> %v\n", h)

	//Output:
	//test: CreateHeaders() -> map[Key:[value] Key2:[value2]]
	//test: CreateHeaders() -> map[Key:[value] Key1:[value1] Key2:[value2]]

}


*/

func Example_SetHeaders() {
	r := httptest.NewRecorder()

	//SetHeaders(r, "key-only")
	//fmt.Printf("test: SetHeaders() [err:%v] [cnt:%v]\n", err, len(r.Result().Header))

	r = httptest.NewRecorder()
	SetHeaders(r, []Attr{{"key1", "val-1"}, {"key-2", "val-2"}})
	fmt.Printf("test: SetHeaders() [cnt:%v]\n", len(r.Result().Header))

	//Output:
	//test: SetHeaders() [cnt:2]

}

func ExampleAddHeader() {
	h := AddHeader(nil, ContentLocation, "https://www.google.com/search")

	fmt.Printf("test: AddHeader() -> [h:%v]\n", h)

	h = AddHeader(h, ContentLocation, "https://www.google.com/search")
	fmt.Printf("test: AddHeader() -> [h:%v]\n", h)

	//Output:
	//test: AddHeader() -> [h:map[Content-Location:[https://www.google.com/search]]]
	//test: AddHeader() -> [h:map[Content-Location:[https://www.google.com/search https://www.google.com/search]]]

}
