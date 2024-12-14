package core

import (
	"fmt"
	"net/http"
)

var (
	iamFunc HttpExchange
)

type testAllow struct{}

func (t testAllow) Exchange(r *http.Request) (*http.Response, *Status) {
	return iamFunc(r)
}

func ExampleRequest() {
	req := NewIAMRequest("github/behavioral-ai/postgres")

	fmt.Printf("test: Request() -> [req:%v] [uri:%v] [from:%v]\n", req != nil, req.URL, req.Header.Get(IAMFrom))

	//Output:
	//test: Request() -> [req:true] [uri:iam:credentials] [from:github/behavioral-ai/postgres]

}

func ExampleResponse_FromUri() {
	resp := NewIAMResponseFromUri("github/behavioral-ai/postgres", StatusOK())
	fmt.Printf("test: Response_FromUri() -> [resp:%v] [statusCode:%v] [uri:%v]\n", resp != nil, resp.StatusCode, resp.Header.Get(IAMUri))

	resp = NewIAMResponseFromUri("github/behavioral-ai/postgres", StatusNotFound())
	fmt.Printf("test: Response_FromUri() -> [resp:%v] [statusCode:%v] [uri:%v]\n", resp != nil, resp.StatusCode, resp.Header.Get(IAMUri))

	//Output:
	//test: Response_FromUri() -> [resp:true] [statusCode:200] [uri:github/behavioral-ai/postgres]
	//test: Response_FromUri() -> [resp:true] [statusCode:404] [uri:github/behavioral-ai/postgres]

}

func ExampleResponse_FromCredentials() {
	resp := NewIAMResponseFromCredentials("user1", "password1", StatusOK())
	fmt.Printf("test: Response_FromCredentials() -> [resp:%v] [statusCode:%v] [user:%v] [pswd:%v]\n", resp != nil, resp.StatusCode, resp.Header.Get(IAMUser), resp.Header.Get(IAMPassword))

	resp = NewIAMResponseFromCredentials("user2", "let-me-in", StatusNotFound())
	fmt.Printf("test: Response_FromCredentials() -> [resp:%v] [statusCode:%v] [user:%v] [pswd:%v]\n", resp != nil, resp.StatusCode, resp.Header.Get(IAMUser), resp.Header.Get(IAMPassword))

	//Output:
	//test: Response_FromCredentials() -> [resp:true] [statusCode:200] [user:user1] [pswd:password1]
	//test: Response_FromCredentials() -> [resp:true] [statusCode:404] [user:user2] [pswd:let-me-in]

}
