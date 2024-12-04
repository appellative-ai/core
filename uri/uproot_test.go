package uri

import (
	"fmt"
	url2 "net/url"
)

func ExampleUproot_Validate() {
	// Empty
	path := ""
	p := Uproot(path)
	fmt.Printf("test: Uproot-Empty(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Path, p.Err)

	// Urn should not be changed
	path = "urn:github.test"
	p = Uproot(path)
	fmt.Printf("test: Uproot-URN(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Path, p.Err)

	// No URN separator, valid with Domain only
	path = "http://localhost:8080/github/advanced-go/search/query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot-Domain-Only(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Path, p.Err)

	// 1 URN separator
	path = "http://localhost:8080/github/advanced-go/search:query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot-Domain+Path(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Path, p.Err)

	// More than 1 URN separator
	path = "http://localhost:8080/github/advanced-go/:search:/query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot->1 URN(%v) -> [ok:%v] [auth:%v] [vers:%v] [path:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Path, p.Err)

	//Output:
	//test: Uproot-Empty() -> [ok:false] [auth:] [vers:] [path:] [err:error: invalid input, URI is empty]
	//test: Uproot-URN(urn:github.test) -> [ok:true] [auth:urn:github.test] [vers:] [path:urn:github.test] [err:<nil>]
	//test: Uproot-Domain-Only(http://localhost:8080/github/advanced-go/search/query?term=golang) -> [ok:true] [auth:github/advanced-go/search/query] [vers:] [path:] [err:<nil>]
	//test: Uproot-Domain+Path(http://localhost:8080/github/advanced-go/search:query?term=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:] [path:query] [err:<nil>]
	//test: Uproot->1 URN(http://localhost:8080/github/advanced-go/:search:/query?term=golang) -> [ok:false] [auth:] [vers:] [path:] [err:error: path has multiple URN separators [/github/advanced-go/:search:/query]]

}

func ExampleUproot() {
	path := "/github/advanced-Go/search:query?term=golang"
	p := Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [rsc:%v] [path:%v] [query:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Resource, p.Path, p.Query, p.Err)

	path = "/github/advanced-go/search:v1/query"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [rsc:%v] [path:%v] [query:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Resource, p.Path, p.Query, p.Err)

	path = "http://localhost:8080/gITHub/advanced-go/search:query?term=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [rsc:%v] [path:%v] [query:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Resource, p.Path, p.Query, p.Err)

	path = "http://localhost:8080/github/advanced-go/search:v1/query"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [rsc:%v] [path:%v] [query:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Resource, p.Path, p.Query, p.Err)

	path = "http://localhost:8080/github/advanced-go/search:v1/query/yahoo?q=golang"
	p = Uproot(path)
	fmt.Printf("test: Uproot(%v) -> [ok:%v] [auth:%v] [vers:%v] [rsc:%v] [path:%v] [query:%v] [err:%v]\n", path, p.Valid, p.Domain, p.Version, p.Resource, p.Path, p.Query, p.Err)

	//Output:
	//test: Uproot(/github/advanced-Go/search:query?term=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:] [rsc:query] [path:query] [query:term=golang] [err:<nil>]
	//test: Uproot(/github/advanced-go/search:v1/query) -> [ok:true] [auth:github/advanced-go/search] [vers:v1] [rsc:query] [path:query] [query:] [err:<nil>]
	//test: Uproot(http://localhost:8080/gITHub/advanced-go/search:query?term=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:] [rsc:query] [path:query] [query:term=golang] [err:<nil>]
	//test: Uproot(http://localhost:8080/github/advanced-go/search:v1/query) -> [ok:true] [auth:github/advanced-go/search] [vers:v1] [rsc:query] [path:query] [query:] [err:<nil>]
	//test: Uproot(http://localhost:8080/github/advanced-go/search:v1/query/yahoo?q=golang) -> [ok:true] [auth:github/advanced-go/search] [vers:v1] [rsc:query] [path:query/yahoo] [query:q=golang] [err:<nil>]

}

/*

	// valid path only and an empty nss
	uri = "/valid-empty-nss?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	// valid embedded path only
	uri = "/github/valid-leading-slash/example-domain/activity:entry"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	// valid URN
	uri = "github.com/valid-no-leading-slash/example-domain/activity:entry"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	uri = "https://www.google.com/valid-uri?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	uri = "https://www.google.com/github.com/valid-uri-nss/search?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)

	uri = "https://www.google.com/github.com/valid-uri-with-nss:search?q=golang"
	nid, nss, ok = Uproot(uri)
	fmt.Printf("test: Uproot(%v) -> [nid:%v] [nss:%v] [ok:%v]\n", uri, nid, nss, ok)


*/

func ExampleUprootDomain() {
	path := "/github/advanced-go/search:yahoo?q=golang"
	url, _ := url2.Parse(path)
	auth := UprootDomain(url)
	fmt.Printf("test: UprootDomain(\"%v\") -> [auth:%v]\n", path, auth)

	path = "github/advanced-go/search:yahoo?q=golang"
	url, _ = url2.Parse(path)
	auth = UprootDomain(url)
	fmt.Printf("test: UprootDomain(\"%v\") -> [auth:%v]\n", path, auth)

	path = "github/adv:anced-go/search:yahoo?q=golang"
	url, _ = url2.Parse(path)
	auth = UprootDomain(url)
	fmt.Printf("test: UprootDomain(\"%v\") -> [auth:%v]\n", path, auth)

	path = "github/advanced-go/searchyahoo?q=golang"
	url, _ = url2.Parse(path)
	auth = UprootDomain(url)
	fmt.Printf("test: UprootDomain(\"%v\") -> [auth:%v]\n", path, auth)

	path = "http://localhost:8080/github.com/advanced-go/example-domain/activity:entry"
	url, _ = url2.Parse(path)
	auth = UprootDomain(url)
	fmt.Printf("test: UprootDomain(\"%v\") -> [auth:%v]\n", path, auth)

	//Output:
	//test: UprootDomain("/github/advanced-go/search:yahoo?q=golang") -> [auth:github/advanced-go/search]
	//test: UprootDomain("github/advanced-go/search:yahoo?q=golang") -> [auth:github/advanced-go/search]
	//test: UprootDomain("github/adv:anced-go/search:yahoo?q=golang") -> [auth:]
	//test: UprootDomain("github/advanced-go/searchyahoo?q=golang") -> [auth:]
	//test: UprootDomain("http://localhost:8080/github.com/advanced-go/example-domain/activity:entry") -> [auth:github.com/advanced-go/example-domain/activity]

}

func _ExampleUrl() {
	uri := "https://user:pswd@www.google.com/search?q=golang"
	u, err := url2.Parse(uri)
	info := u.User

	fmt.Printf("test: Url(\"%v\") -> [%v] [err:%v] [%v]\n", uri, u, err, info)

	uri = "https://www.google.com/search?q=golang"
	u, err = url2.Parse(uri)
	info = u.User

	fmt.Printf("test: Url(\"%v\") -> [%v] [err:%v] [%v]\n", uri, u, err, info)

	//Output:
	//fail

}
