package host

import (
	"fmt"
)

var (
	networkConfig = "file://[cwd]/resource/network-config.json"
)

func ExampleReadConfig() {
	net, err := readConfig(networkConfig)
	fmt.Printf("test: readConfig(\"%v\") -> [map:%v] [err:%v]\n", networkConfig, len(net), err)

	//Output:
	//test: readConfig("file://[cwd]/resource/network-config.json") -> [map:6] [err:<nil>]

}

func ExampleParseConfig() {
	s := "app-host=localhost:8082,log=true,invalid2=,route-name=test-route,timeout=2m,invalid="
	m, err := parseConfig(s)
	fmt.Printf("test: parseConfig(\"%v\") -> %v [err:%v]\n", s, m, err)

	//Output:
	//test: parseConfig("app-host=localhost:8082,log=true,invalid2=,route-name=test-route,timeout=2m,invalid=") -> map[app-host:localhost:8082 log:true route-name:test-route timeout:2m] [err:<nil>]

}

func ExampleNewPath() {

	fmt.Printf("test: newPath(\"%v\") -> %v\n", networkConfig, newPath(networkConfig, "@path=cache-config.json"))

	//Output:
	//test: newPath("file://[cwd]/resource/network-config.json") -> file://[cwd]/resource/cache-config.json

}

func ExampleDefineNetwork() {

	net, err := DefineNetwork(networkConfig, []string{"redirect", "rate-limiter", "cache", "routing", "logging", "authorization"})
	fmt.Printf("test: DefineNetwork() -> %v [err:%v]\n", net, err)

	//Output:
	//test: DefineNetwork() -> [count:4] [{test:resiliency:agent/rate-limiting/request/http map[load-size:567 off-peak-duration:5m peak-duration:750ms rate-burst:12 rate-limit:1234]} {test:resiliency:agent/redirect/request/http map[interval:4m new-path:/resource/v2 original-path:/resource/v1 percentile-threshold:99/1500ms rate-burst:12 rate-limit:1234 status-code-threshold:10]} {test:resiliency:agent/cache/request/http map[cache-control:no-store, no-cache, max-age=0 fri:22-23 host:www.google.com interval:4m mon:8-16 sat:3-8 sun:13-15 thu:0-23 timeout:750ms tue:6-10 wed:12-12]} {test:resiliency:agent/routing/request/http map[app-host:localhost:8082 log:true route-name:test-route timeout:2m]}] [err:<nil>]

}
