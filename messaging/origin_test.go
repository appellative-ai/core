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
	//test: NewOrigin() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone sub-zone host service-name instance-id collective domain}] [status:OK]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region  sub-zone host service-name instance-id collective domain}]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone  host service-name instance-id collective domain}]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone sub-zone  service-name instance-id collective domain}]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone sub-zone host service-name instance-id collective domain}]

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
	//test: NewOrigin() -> [{        }] [status:Invalid Content [config map does not contain key: collective]]
	//test: NewOrigin() -> [{       collective }] [status:Invalid Content [config map does not contain key: domain]]
	//test: NewOrigin() -> [{       collective domain}] [status:Invalid Content [config map does not contain key: region]]
	//test: NewOrigin() -> [{ region      collective domain}] [status:Invalid Content [config map does not contain key: zone]]
	//test: NewOrigin() -> [{ region zone     collective domain}] [status:Invalid Content [config map does not contain key: host]]
	//test: NewOrigin() -> [{collective:domain:service/region/zone/host region zone  host host  collective domain}] [status:OK]

}

func ExampleIsLocalCollectiveNewOrigin() {
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
	o, _ := NewOrigin(m)

	name1 := ""
	fmt.Printf("test: IsLocalCollective(\"%v\") -> [local:%v]\n", name1, o.IsLocalCollective(name1))

	name1 = o.Collective
	fmt.Printf("test: IsLocalCollective(\"%v\") -> [local:%v]\n", name1, o.IsLocalCollective(name1))

	name1 = o.Collective + ":"
	fmt.Printf("test: IsLocalCollective(\"%v\") -> [local:%v]\n", name1, o.IsLocalCollective(name1))

	//Output:
	//test: IsLocalCollective("") -> [local:false]
	//test: IsLocalCollective("collective") -> [local:false]
	//test: IsLocalCollective("collective:") -> [local:true]

}
