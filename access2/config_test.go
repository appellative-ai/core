package access2

import (
	"fmt"
	"github.com/behavioral-ai/core/iox"
)

const (
	opsPath = "file://[cwd]/accesstest/logging-operators.json"
)

func readFunc() ([]byte, error) {
	return iox.ReadFile(opsPath)
}

func ExampleLoadOperators() {
	err := LoadOperators(readFunc)

	fmt.Printf("test: LoadOperators() -> [err:%v] [ops:%v]\n", err, defaultOperators)

	//Output:
	//test: LoadOperators() -> [err:<nil>] [ops:[{start-time %START_TIME%} {duration-ms %DURATION%} {traffic %TRAFFIC%} {route %ROUTE%} {method %METHOD%} {host %HOST%} {path %PATH%} {status-code %STATUS_CODE%} {timeout-ms %TIMEOUT_DURATION%} {rate-limit %RATE_LIMIT%} {redirect %REDIRECT%}]]

}
