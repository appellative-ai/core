package jsonx

import "fmt"

func ExampleNewStatusFrom_Const() {
	status := NewStatusFrom("")
	fmt.Printf("test: NewStatusFrom(nil) -> [code:%v]\n", status)

	uri := StatusOKUri
	status = NewStatusFrom(uri)
	fmt.Printf("test: NewStatusFrom(\"%v\") -> [code:%v]\n", uri, status)

	uri = StatusNotFoundUri
	status = NewStatusFrom(uri)
	fmt.Printf("test: NewStatusFrom(\"%v\") -> [code:%v] [status:%v]\n", uri, status, status)

	uri = StatusTimeoutUri
	status = NewStatusFrom(uri)
	fmt.Printf("test: NewStatusFrom(\"%v\") -> [code:%v] [status:%v]\n", uri, status, status)

	//Output:
	//test: NewStatusFrom(nil) -> [code:200]
	//test: NewStatusFrom("urn:status:ok") -> [code:200]
	//test: NewStatusFrom("urn:status:notfound") -> [code:404] [status:Not Found]
	//test: NewStatusFrom("urn:status:timeout") -> [code:504] [status:Timeout]

}
