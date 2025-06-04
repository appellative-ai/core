package messaging

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Line1 string
	Line2 string
	City  string
	State string
	Zip   string
}

func ExampleBinary() {
	bytes := []byte("this is a test buffer")
	t, status := Unmarshal[[]byte](&Content{Type: ContentTypeBinary, Value: bytes})
	fmt.Printf("test: Unmarshal[[]byte]() -> [%v] [status:%v]\n", string(t), status)

	/*
		ct, status2 := Resolver.Representation(name, fragment)
		fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

		if buf, ok := ct.Value.([]byte); ok {
			fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", string(buf), status2)
		}

		s3, status3 := Resolve[[]byte](name, fragment, Resolver)
		fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", string(s3), status3)

	*/

	//Output:
	//test: Unmarshal[[]byte]() -> [this is a test buffer] [status:OK]

}

func ExampleString() {
	s := "this is a test string"

	t, status := Unmarshal[string](&Content{Type: ContentTypeText, Value: []byte(s)})
	fmt.Printf("test: Unmarshal[string]() -> [%v] [status:%v]\n", t, status)

	/*
		ct, status2 := Resolver.Representation(name, fragment)
		fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

		if buf, ok := ct.Value.([]byte); ok {
			fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", string(buf), status2)
		}

		s3, status3 := Resolve[string](name, fragment, Resolver)
		fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", string(s3), status3)


	*/

	//Output:
	//test: Unmarshal[string]() -> [this is a test string] [status:OK]

}

func ExampleResolveType() {
	addr := Address{
		Line1: "123 Main",
		Line2: "",
		City:  "Anytown",
		State: "Ohio",
		Zip:   "54321",
	}
	buf, err := json.Marshal(&addr)
	if err != nil {
		fmt.Printf("test: json.Marshal() -> [err:%v]\n", err)
	}
	t, status := Unmarshal[Address](&Content{Type: ContentTypeJson, Value: buf})
	fmt.Printf("test: Unmarshal[Address]() -> [%v] [status:%v]\n", t, status)

	/*
		ct, status2 := Resolver.Representation(name, fragment)
		fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

		if buf, ok := ct.Value.([]byte); ok {
			fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
		}

		s3, status3 := Resolve[Address](name, fragment, Resolver)
		fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", s3, status3)


	*/

	//Output:
	//test: Unmarshal[Address]() -> [{123 Main  Anytown Ohio 54321}] [status:OK]

}

/*
func ExampleResolveMap() {
	NewAgent()
	m := map[string]string{
		"Line1": "123 Main",
		"Line2": "",
		"City":  "Anytown",
		"State": "Ohio",
		"Zip":   "54321",
	}
	name := "core:type/map"
	fragment := "v2"

	status := Resolver.AddRepresentation(name, fragment, "author", m)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name, fragment)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	s3, status3 := Resolve[map[string]string](name, fragment, Resolver)
	fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", s3, status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v2 type: application/json value: true] [status:OK]
	//test: Representation() -> [value:77] [status:OK]
	//test: Resolve() -> [value:map[City:Anytown Line1:123 Main Line2: State:Ohio Zip:54321]] [status:OK]

}

*/
