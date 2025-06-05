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
		//"Collective: "collective",
		//Domain:     "domain",
	}
	o, status := NewOrigin(m, "collective", "domain")
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
	//test: NewOrigin() -> [collective:domain:service/region/zone/sub-zone/service-name] [status:<nil>]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]
	//test: Name() -> [collective:domain:service/region/zone/sub-zone/service-name]

}

/*
func _ExampleOrigin_Uri() {
	target := originT{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "instance-id",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target.Uri("class"))

	target.SubZone = ""
	fmt.Printf("test: Origin_Uri_No_SubZone()    -> [%v]\n", target.Uri("class"))

	//Output:
	//test: Origin_Uri_SubZone()       -> [class:region.zone.sub-zone.host]
	//test: Origin_Uri_No_SubZone()    -> [class:region.zone.host]

}

func _ExampleOrigin_String() {
	target := originT{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "instance-id",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target)

	target.SubZone = ""
	fmt.Printf("test: Origin_Uri_No_SubZone()    -> [%v]\n", target)

	//Output:
	//test: Origin_Uri_SubZone()       -> [region.zone.sub-zone.host]
	//test: Origin_Uri_No_SubZone()    -> [region.zone.host]

}


*/
