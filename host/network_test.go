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

	fmt.Printf("test: newPath(\"%v\") -> %v\n", networkConfig, newPath(networkConfig, "cache-config.json"))

	//Output:
	//test: newPath("file://[cwd]/resource/network-config.json") -> file://[cwd]/resource/cache-config.json

}

func ExampleDefineNetwork() {
	redirect := "redirect"
	limiter := "rate-limiter"
	cache := "cache"
	routing := "routing"
	logging := "logging"
	authz := "authorization"
	net, _ := DefineNetwork(networkConfig, []string{redirect, limiter, cache, routing, logging, authz})

	fmt.Printf("test: Resource(\"%v\") -> %v\n", redirect, net.Load(redirect))
	fmt.Printf("test: Resource(\"%v\") -> %v\n", limiter, net.Load(limiter))
	fmt.Printf("test: Resource(\"%v\") -> %v\n", cache, net.Load(cache))
	fmt.Printf("test: Resource(\"%v\") -> %v\n", routing, net.Load(routing))
	fmt.Printf("test: Resource(\"%v\") -> %v\n", logging, net.Load(logging))
	fmt.Printf("test: Resource(\"%v\") -> %v\n", authz, net.Load(authz))

	//Output:
	//test: Resource("redirect") -> {redirect test:resiliency:agent/redirect/request/http <nil> map[interval:4m new-path:/resource/v2 original-path:resource/v1 percentile-threshold:99/1500ms rate-burst:12 rate-limit:1234 status-code-threshold:10]}
	//test: Resource("rate-limiter") -> {rate-limiter test:resiliency:agent/rate-limiting/request/http <nil> map[load-size:567 off-peak-duration:5m peak-duration:750ms rate-burst:12 rate-limit:1234]}
	//test: Resource("cache") -> {cache test:resiliency:agent/cache/request/http <nil> map[cache-control:no-store, no-cache, max-age=0 fri:22-23 host:www.google.com interval:4m mon:8-16 sat:3-8 sun:13-15 thu:0-23 timeout:750ms tue:6-10 wed:12-12]}
	//test: Resource("routing") -> {routing test:resiliency:agent/routing/request/http <nil> map[app-host:localhost:8082 log:true route-name:test-route timeout:2m]}
	//test: Resource("logging") -> {logging test:resiliency:link/logging/access <nil> map[]}
	//test: Resource("authorization") -> {authorization test:resiliency:link/authorization/http <nil> map[]}

}
