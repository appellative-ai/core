package rest

import (
	"fmt"
)

func ExampleAuthorization_Chain() {
	//name := "agent/authorization"
	chain := BuildExchangeChain([]any{Authorization})
	fmt.Printf("test: BuildExchangeChain() -> %v\n", chain != nil)

	//repository.RegisterExchangeLink(name, Authorization)
	//chain = rest.BuildExchangeChain([]any{repository.ExchangeLink(name)})
	//fmt.Printf("test: repository.ExchangeLink() -> %v\n", chain != nil)

	//Output:
	//fail

}
