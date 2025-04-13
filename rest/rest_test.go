package rest

import (
	"context"
	"fmt"
	"net/http"
)

func ExampleBuildChain_Link() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1LinkFn, do2LinkFn, do3LinkFn)
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response
	//test: Do1-Link() -> response

}

func ExampleBuildChain_Chainable() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, do2{}, do3{})
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response
	//test: Do1-Link() -> response

}

func ExampleBuildChain_Any() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, do2LinkFn, do3{}, do4LinkFn)
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do4-Link() -> request
	//test: Do4-Link() -> response
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response
	//test: Do1-Link() -> response

}

func ExampleBuildChain_Abbreviated() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1LinkFn, do2LinkFn, do3FailLinkFn, do4LinkFn)
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Fail-Link() -> request
	//test: Do3-Fail-Link() -> response
	//test: Do2-Link() -> response
	//test: Do1-Link() -> response

}

func ExampleBuildChain_Exchange() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1LinkFn, do2LinkFn, do3LinkFn, do5ExchangeFn)
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do5-Exchange() -> request
	//test: Do5-Exchange() -> response
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response
	//test: Do1-Link() -> response

}

func ExampleBuildChain_Exchangeable() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1LinkFn, do2LinkFn, do3LinkFn, do5{})
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do5-Exchange() -> request
	//test: Do5-Exchange() -> response
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response
	//test: Do1-Link() -> response

}

func do1LinkFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do1-Link() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do1-Link() -> response\n")
		return
	}
}

func do2LinkFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do2-Link() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do2-Link() -> response\n")
		return
	}
}

func do3LinkFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do3-Link() -> request\n")
		//fmt.Printf("test: Do3() -> response\n")
		//return &http.Response{StatusCode: http.StatusBadRequest}, nil
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do3-Link() -> response\n")
		return
	}
}

func do3FailLinkFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do3-Fail-Link() -> request\n")
		fmt.Printf("test: Do3-Fail-Link() -> response\n")
		return &http.Response{StatusCode: http.StatusBadRequest}, nil
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do3-Fail-Link() -> response\n")
		return
	}
}

func do4LinkFn(next Exchange) Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		fmt.Printf("test: Do4-Link() -> request\n")
		if next != nil {
			resp, err = next(req)
		} else {
			resp = &http.Response{StatusCode: http.StatusOK}
		}
		fmt.Printf("test: Do4-Link() -> response\n")
		return
	}
}

func do5ExchangeFn(req *http.Request) (resp *http.Response, err error) {
	fmt.Printf("test: Do5-Exchange() -> request\n")
	resp = &http.Response{StatusCode: http.StatusOK}
	fmt.Printf("test: Do5-Exchange() -> response\n")
	return
}

type do1 struct{}

func (d do1) Link(next Exchange) Exchange {
	return do1LinkFn(next)
}

type do2 struct{}

func (d do2) Link(next Exchange) Exchange {
	return do2LinkFn(next)
}

type do3 struct{}

func (d do3) Link(next Exchange) Exchange {
	return do3LinkFn(next)
}

type do3Fail struct{}

func (d do3Fail) Link(next Exchange) Exchange {
	return do3FailLinkFn(next)
}

type do4 struct{}

func (d do4) Link(next Exchange) Exchange {
	return do4LinkFn(next)
}

type do5 struct{}

func (d do5) Exchange(r *http.Request) (*http.Response, error) {
	return do5ExchangeFn(r)
}

func _ExampleBuildChain_Panic_Type() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, req, do3{}, do4LinkFn)
	ex(req)

	//Output:
	//fail
}

func _ExampleBuildChain_Panic_Nil() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1{}, nil, do3{}, do4LinkFn)
	ex(req)

	//Output:
	//fail
}

func _ExampleBuildChain_Exchange_Panic() {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	ex := BuildChain(do1LinkFn, do2LinkFn, do5ExchangeFn, do3LinkFn)
	ex(req)

	//Output:
	//test: Do1-Link() -> request
	//test: Do2-Link() -> request
	//test: Do3-Link() -> request
	//test: Do5-Exchange() -> request
	//test: Do5-Exchange() -> response
	//test: Do3-Link() -> response
	//test: Do2-Link() -> response

}
