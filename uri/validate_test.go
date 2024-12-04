package uri

import (
	"fmt"
	"net/http"
)

func ExampleValidateURL_Invalid() {
	_, status := ValidateURL(nil, "")
	fmt.Printf("test: ValidateURL(nil,\"\") -> [%v]\n", status)

	path := "test"
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, status = ValidateURL(req.URL, path)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [%v]\n", req.URL.Path, path, status)

	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, status = ValidateURL(req.URL, "")
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	_, status = ValidateURL(req.URL, path)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github/advanced-go/http2", nil)
	_, status = ValidateURL(req.URL, path)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [%v]\n", req.URL.Path, path, status)

	//Output:
	//test: ValidateURL(nil,"") -> [error: URL is nil]
	//test: ValidateURL("","test") -> [error: invalid input, URI is empty]
	//test: ValidateURL("","test") -> [error: domain is empty]
	//test: ValidateURL("/search","github/advanced-go/http2") -> [error: invalid URI, domain does not match: "/search" "github/advanced-go/http2"]
	//test: ValidateURL("/github/advanced-go/http2","github/advanced-go/http2") -> [error: invalid URI, path only contains a domain: "/github/advanced-go/http2"]

}

func ExampleValidateRequest() {
	auth := "github/advanced-go/httpx"
	rscSearch := ":search?q=golang"
	uri := "https://www.google.com/" + auth + rscSearch
	rscVerSearch := ":v1/search/yahoo?q=golang"
	uri2 := "https://www.google.com/" + auth + rscVerSearch

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	p, status := ValidateURL(req.URL, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [%v] [ver:%v] [rsc:%v] [path:%v] [query:%v]\n", uri, auth, status, p.Version, p.Resource, p.Path, p.Query)

	req, _ = http.NewRequest(http.MethodGet, uri2, nil)
	p, status = ValidateURL(req.URL, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [%v] [ver:%v] [rsc:%v] [path:%v] [query:%v]\n", uri2, auth, status, p.Version, p.Resource, p.Path, p.Query)

	//Output:
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:search?q=golang","github/advanced-go/httpx") -> [<nil>] [ver:] [rsc:search] [path:search] [query:q=golang]
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:v1/search/yahoo?q=golang","github/advanced-go/httpx") -> [<nil>] [ver:v1] [rsc:search] [path:search/yahoo] [query:q=golang]

}

func ExampleValidateURL_Domain() {
	auth := "github/advanced-go/stdlib"
	req, _ := http.NewRequest(http.MethodGet, DomainRootPath, nil)
	p, status := ValidateURL(req.URL, auth)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [%v] [ver:%v] [path:%v]\n", DomainPath, auth, status, p.Version, p.Path)

	//Output:
	//test: ValidateURL("domain","github/advanced-go/stdlib") -> [<nil>] [ver:] [path:domain]

}
