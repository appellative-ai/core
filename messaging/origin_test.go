package messaging

import "fmt"

func ExampleNewOrigin() {
	m := map[string]string{
		RegionKey:      "region",
		ZoneKey:        "zone",
		SubZoneKey:     "sub-zone",
		HostKey:        "host",
		ServiceNameKey: "service-name",
		//InstanceId: "instance-id",
		CollectiveKey: "collective",
		DomainKey:     "domain",
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
	//test: NewOrigin() -> [collective:domain:service/region/zone/sub-zone/service-name] [status:OK]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]

}
