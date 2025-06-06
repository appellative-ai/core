package messaging

import "fmt"

func ExampleNewOrigin() {
	m := map[string]string{
		RegionKey:      "region",
		ZoneKey:        "zone",
		SubZoneKey:     "sub-zone",
		HostKey:        "host",
		ServiceNameKey: "service-name",
		InstanceIdKey:  "instance-id",
		CollectiveKey:  "collective",
		DomainKey:      "domain",
	}
	o, status := NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	o.Zone = ""
	fmt.Printf("test: Name() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = ""
	fmt.Printf("test: Name() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = ""
	fmt.Printf("test: Name() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = "host"
	o.InstanceId = "instance-id"
	fmt.Printf("test: Name() -> [%v]\n", o)

	//Output:
	//test: NewOrigin() -> [collective:domain:service/region/zone/sub-zone/service-name#instance-id] [status:OK]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name#instance-id]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name#instance-id]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name#instance-id]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name#instance-id]

}

func ExampleNewOrigin_Error() {
	m := make(map[string]string)

	o, status := NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	m[CollectiveKey] = "collective"
	o, status = NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	m[DomainKey] = "domain"
	o, status = NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	m[RegionKey] = "region"
	o, status = NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	m[ZoneKey] = "zone"
	o, status = NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	m[HostKey] = "host"
	o, status = NewOrigin(m)
	fmt.Printf("test: NewOrigin() -> [%v] [status:%v]\n", o, status)

	//Output:
	//test: NewOrigin() -> [] [status:Invalid Content [config map does not contain key: collective]]
	//test: NewOrigin() -> [] [status:Invalid Content [config map does not contain key: domain]]
	//test: NewOrigin() -> [] [status:Invalid Content [config map does not contain key: region]]
	//test: NewOrigin() -> [] [status:Invalid Content [config map does not contain key: zone]]
	//test: NewOrigin() -> [] [status:Invalid Content [config map does not contain key: host]]
	//test: NewOrigin() -> [collective:domain:service/region/zone/host] [status:OK]

}
