package messaging

import "fmt"

func ExampleNewOrigin() {
	o := Origin{
		Region:  "region",
		Zone:    "zone",
		SubZone: "sub-zone",
		Host:    "host",
		//InstanceId: "instance-id",
	}
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = ""
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = ""
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = ""
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = "host"
	o.InstanceId = "instance-id"
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	//Output:
	//test: NewOrigin() -> [core:service/region/zone/sub-zone/host]
	//test: NewOrigin() -> [core:service/region/sub-zone/host]
	//test: NewOrigin() -> [core:service/region/zone/host]
	//test: NewOrigin() -> [core:service/region/zone/sub-zone]
	//test: NewOrigin() -> [core:service/region/zone/sub-zone/host#instance-id]

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
