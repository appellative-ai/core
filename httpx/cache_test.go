package httpx

import (
	"fmt"
	"net/http"
	"time"
)

var (
	cache = newCache()
)

func ExampleNewCache() {
	c := NewResponseCache()
	uri := "https://www.google.com/search?q=golang"
	s := "this is string content"

	h := make(http.Header)
	h.Set("key-1", "value-1")
	h.Set("key-2", "value-2")
	h.Set("key-3", "value-3")
	resp := NewResponse(http.StatusOK, h, s)
	err1 := c.Put(uri, resp)
	if err1 != nil {
		fmt.Printf("test: cache.Put() -> [err:%v]\n", err1)
	}

	req2, _ := http.NewRequest(http.MethodGet, uri, nil)
	resp1 := c.Get(req2.URL.String())
	buf, err := readAll(resp1.Body)
	if err != nil {
		fmt.Printf("test: readAll() -> [buf:%v] [err:%v]\n", buf, err)
	}
	fmt.Printf("test: NewCache(\"%v\") -> [%v] [%v] [%v]\n", uri, resp.StatusCode, resp.Header, string(buf))

	uri = "https://bing.com/search?q=golang"
	//req3, _ := http.NewRequest(http.MethodGet,uri, nil)
	resp = c.Get(uri)
	buf, err = readAll(resp.Body)
	fmt.Printf("test: NewCache(\"%v\") -> [%v] [%v] [%v] [err:%v]\n", uri, resp.StatusCode, resp.Header, buf, err)

	uri = "https://www.google.com/search?q=golang"
	req2, _ = http.NewRequest(http.MethodGet, uri, nil)
	resp1 = c.Get(req2.URL.String())
	buf, err = readAll(resp1.Body)
	if err != nil {
		fmt.Printf("test: readAll() -> [buf:%v] [err:%v]\n", buf, err)
	}
	fmt.Printf("test: NewCache(\"%v\") -> [%v] [%v] [%v]\n", uri, resp1.StatusCode, resp1.Header, string(buf))

	//Output:
	//test: NewCache("https://www.google.com/search?q=golang") -> [200] [map[Key-1:[value-1] Key-2:[value-2] Key-3:[value-3]]] [this is string content]
	//test: NewCache("https://bing.com/search?q=golang") -> [404] [map[]] [[]] [err:<nil>]
	//test: NewCache("https://www.google.com/search?q=golang") -> [200] [map[Key-1:[value-1] Key-2:[value-2] Key-3:[value-3]]] [this is string content]

}

func putCache(url string, timeout time.Duration) (*http.Response, error) {
	// create request and process exchange
	ctx, cancel := NewContext(nil, timeout)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(AcceptEncoding, GzipEncoding)
	resp, err1 := ExchangeWithTimeout(timeout, nil)(req)
	if err1 != nil {
		return resp, err1
	}

	cache.Put(url, resp)
	return resp, nil
}

func ExampleCache_No_Timeout() {
	url := "https://www.google.com/search?q=golang"
	timeout := time.Millisecond * 0
	fmt.Printf("test: ExampleCache() [url:%v] [timeout:%v]\n", url, timeout)

	resp, err := putCache(url, timeout)
	fmt.Printf("test: cache.put() [status:%v] [%v]\n", resp.StatusCode, err)

	// Get cached response
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(AcceptEncoding, GzipEncoding)
	resp1 := cache.Get(req.URL.String())
	fmt.Printf("test: cache.get() [status:%v] [header:%v]\n", resp1.StatusCode, resp.Header != nil)

	// verify that the response body can be read
	if resp1.StatusCode == http.StatusOK {
		buf, err1 := readAll(resp1.Body)
		fmt.Printf("test: readAll() [err:%v] [buf:%v]\n", err1, len(buf))
	}

	//Output:
	//test: ExampleCache() [url:https://www.google.com/search?q=golang] [timeout:0s]
	//test: cache.Put() [status:200] [<nil>]
	//test: cache.Get() [status:200] [header:true]
	//test: readAll() [err:<nil>] [buf:40984]

}

func ExampleCache_Timeout_504() {
	url := "https://www.google.com/search?q=erlang"
	timeout := time.Millisecond * 10
	fmt.Printf("test: ExampleCache() [url:%v] [timeout:%v]\n", url, timeout)

	resp, err := putCache(url, timeout)
	fmt.Printf("test: cache.put() [status:%v] [%v]\n", resp.StatusCode, err)

	// Get cached response
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(AcceptEncoding, GzipEncoding)
	resp = cache.Get(req.URL.String())
	fmt.Printf("test: cache.get() [status:%v]\n", resp.StatusCode)

	// verify that the response body can be read
	if resp.StatusCode == http.StatusOK {
		buf, err1 := readAll(resp.Body)
		fmt.Printf("test: readAll() [err:%v] [buf:%v]\n", err1, len(buf))
	}

	//Output:
	//test: ExampleCache() [url:https://www.google.com/search?q=erlang] [timeout:10ms]
	//test: cache.put() [status:504] [Get "https://www.google.com/search?q=erlang": context deadline exceeded]
	//test: cache.get() [status:404]

}

func ExampleCache_Timeout_200() {
	url := "https://www.google.com/search?q=pascal"
	timeout := time.Second * 5
	fmt.Printf("test: ExampleCache() [url:%v] [timeout:%v]\n", url, timeout)

	resp, err := putCache(url, timeout)
	fmt.Printf("test: cache.put() [status:%v] [%v]\n", resp.StatusCode, err)

	// Get cached response
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(AcceptEncoding, GzipEncoding)
	resp1 := cache.Get(req.URL.String())
	fmt.Printf("test: cache.get() [status:%v] [header:%v]\n", resp.StatusCode, resp.Header != nil)

	// verify that the response body can be read
	if resp1.StatusCode == http.StatusOK {
		buf, err1 := readAll(resp1.Body)
		fmt.Printf("test: readAll() [err:%v] [buf:%v]\n", err1, len(buf))
	}

	//Output:
	//test: ExampleCache() [url:https://www.google.com/search?q=pascal] [timeout:5s]
	//test: cache.put() [status:200] [<nil>]
	//test: cache.get() [status:200] [header:true]
	//test: readAll() [err:<nil>] [buf:40912]

}
