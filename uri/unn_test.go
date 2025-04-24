package uri

import (
	"fmt"
)

func ExampleParseUnn() {
	uri := "unn:behavioral-ai.github.com:resiliency:agent/collective/namespace"
	u := ParseUnn(uri)
	fmt.Printf("test: ParseUnn(\"%v\") -> [%v] [%v] [%v] [%v] [%v] [%v] [err:%v]\n", uri, u.Authority, u.Domain, u.Kind, u.Path, u.Resource, u.Fragment, u.Err)

	uri = "unn:behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5"
	u = ParseUnn(uri)
	fmt.Printf("test: ParseUnn(\"%v\") -> [%v] [%v] [%v] [%v] [%v] [%v] [err:%v]\n", uri, u.Authority, u.Domain, u.Kind, u.Path, u.Resource, u.Fragment, u.Err)

	uri = "https://somedomain.com/behavioral-ai.github.com:resiliency:agent/collective/namespace:advice#2.3.5"
	u = ParseUnn(uri)
	fmt.Printf("test: ParseUnn(\"%v\") -> [%v] [%v] [%v] [%v] [%v] [%v] [err:%v]\n", uri, u.Authority, u.Domain, u.Kind, u.Path, u.Resource, u.Fragment, u.Err)

	//Output:
	//test: ParseUnn("unn:behavioral-ai.github.com:resiliency:agent/collective/namespace") -> [behavioral-ai.github.com] [resiliency] [agent] [collective/namespace] [] [] [err:<nil>]
	//test: ParseUnn("unn:behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5") -> [behavioral-ai.github.com] [resiliency] [agent] [collective/namespace] [state] [1.3.5] [err:<nil>]
	//test: ParseUnn("https://somedomain.com/behavioral-ai.github.com:resiliency:agent/collective/namespace:advice#2.3.5") -> [behavioral-ai.github.com] [resiliency] [agent] [collective/namespace] [advice] [2.3.5] [err:<nil>]

}

func ExampleBuildUnn() {
	authority := "behavioral-ai.github.com"
	domain := "resiliency"
	kind := "agent"
	path := "collective/namespace"
	resource := "state"
	fragment := "1.3.5"

	uri := BuildUnnFrom(authority, domain, kind, path, "", "")
	fmt.Printf("test: BuildUnnFrom() -> [%v]\n", uri)

	uri = BuildUnnFrom(authority, domain, kind, path, resource, "")
	fmt.Printf("test: BuildUnnFrom() -> [%v]\n", uri)

	uri = BuildUnnFrom(authority, domain, kind, path, "", fragment)
	fmt.Printf("test: BuildUnnFrom() -> [%v]\n", uri)

	uri = BuildUnnFrom(authority, domain, kind, path, resource, fragment)
	fmt.Printf("test: BuildUnnFrom() -> [%v]\n", uri)

	uri = BuildUnn(&Unn{
		Authority: authority,
		Domain:    domain,
		Kind:      kind,
		Path:      path,
		Resource:  resource,
		Fragment:  fragment,
		Err:       nil,
	})
	fmt.Printf("test: BuildUnn() -> [%v]\n", uri)

	//Output:
	//test: BuildUnnFrom() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/namespace]
	//test: BuildUnnFrom() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/namespace:state]
	//test: BuildUnnFrom() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/namespace#1.3.5]
	//test: BuildUnnFrom() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5]
	//test: BuildUnn() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5]

}

func ExampleParseKind() {
	unn := new(Unn)
	s := "agentcollective-namespace"
	err := parseKind(s, unn)
	fmt.Printf("test: parseKind(\"%v\") -> [kind:%v] [path:%v] [err:%v]\n", s, unn.Kind, unn.Path, err)

	s = "agent/collective/namespace"
	err = parseKind(s, unn)
	fmt.Printf("test: parseKind(\"%v\") -> [kind:%v] [path:%v] [err:%v]\n", s, unn.Kind, unn.Path, err)

	//Output:
	//test: parseKind("agentcollective-namespace") -> [kind:] [path:] [err:invalid argument: no path for agent [agentcollective-namespace]]
	//test: parseKind("agent/collective/namespace") -> [kind:agent] [path:collective/namespace] [err:<nil>]

}

func ExampleParseResource() {
	unn := new(Unn)
	s := "state"
	parseResource(s, unn)
	fmt.Printf("test: parseResource(\"%v\") -> [resource:%v] [fragment:%v]\n", s, unn.Resource, unn.Fragment)

	s = "state#1.2.3"
	parseResource(s, unn)
	fmt.Printf("test: parseResource(\"%v\") -> [resource:%v] [fragment:%v]\n", s, unn.Resource, unn.Fragment)

	//Output:
	//test: parseResource("state") -> [resource:state] [fragment:]
	//test: parseResource("state#1.2.3") -> [resource:state] [fragment:1.2.3]

}

func ExampleUprootUnn() {
	s := "state"
	uri := uprootUnn(s)
	fmt.Printf("test: uprootUnn(\"%v\") -> [%v]\n", s, uri)

	s = "unnstate"
	uri = uprootUnn(s)
	fmt.Printf("test: uprootUnn(\"%v\") -> [%v]\n", s, uri)

	s = "unn:state"
	uri = uprootUnn(s)
	fmt.Printf("test: uprootUnn(\"%v\") -> [%v]\n", s, uri)

	//Output:
	//test: uprootUnn("state") -> [state]
	//test: uprootUnn("unnstate") -> [unnstate]
	//test: uprootUnn("unn:state") -> [state]

}

func ExampleUprootURL() {
	s := "https://www.google.com/search?q=golang"
	uri := uprootUnn(s)
	fmt.Printf("test: uprootURL(\"%v\") -> [%v]\n", s, uri)

	s = "https://somedomain.com/behavioral-ai.github.com:resiliency:agent/collective/namespace"
	uri = uprootUnn(s)
	fmt.Printf("test: uprootURL(\"%v\") -> [%v]\n", s, uri)

	s = "https://somedomain.com/behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5"
	uri = uprootUnn(s)
	fmt.Printf("test: uprootUnn(\"%v\") -> [%v]\n", s, uri)

	//Output:
	//test: uprootURL("https://www.google.com/search?q=golang") -> [search]
	//test: uprootURL("https://somedomain.com/behavioral-ai.github.com:resiliency:agent/collective/namespace") -> [behavioral-ai.github.com:resiliency:agent/collective/namespace]
	//test: uprootUnn("https://somedomain.com/behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5") -> [behavioral-ai.github.com:resiliency:agent/collective/namespace:state#1.3.5]

}
