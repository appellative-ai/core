package access2

import "fmt"

func ExampleNewValues() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	values := NewValues(o)
	fmt.Printf("test: NewValues() -> [%v]\n", values)

	//Output:
	//test: NewValues() -> [map[host:[host] region:[region] sub-zone:[sub-zone] zone:[zone]]]

}

func ExampleNewOrigin() {
	o := Origin{
		Region:  "region",
		Zone:    "zone",
		SubZone: "sub-zone",
		Host:    "host",
		//Route:      "route",
		InstanceId: "",
	}
	values := NewValues(o)
	o = NewOrigin(values)
	fmt.Printf("test: NewOrigin() -> [%v]\n", o)

	//Output:
	//test: NewOrigin() -> [{region zone sub-zone host }]

}

func ExampleOrigin_Tag() {
	o := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	o.Zone = ""
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	o.Host = ""
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	o.SubZone = ""
	fmt.Printf("test: Tag() -> [%v]\n", o.Tag())

	//Output:
	//test: Tag() -> [region:zone:sub-zone:host]
	//test: Tag() -> [region:sub-zone:host]
	//test: Tag() -> [region:sub-zone]
	//test: Tag() -> [region]

}

func ExampleOrigin_Uri() {
	target := Origin{
		Region:     "region",
		Zone:       "zone",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}

	fmt.Printf("test: Origin_Uri_SubZone()       -> [%v]\n", target.Uri("class"))
	//target.Route = "route"
	//fmt.Printf("test: Origin_Uri_SubZone_Route() -> [%v]\n", target.Uri("class"))

	target.SubZone = ""
	//target.Route = ""
	fmt.Printf("test: Origin_Uri_No_SubZone()    -> [%v]\n", target.Uri("class"))

	//Output:
	//test: Origin_Uri_SubZone()       -> [class:region.zone.sub-zone.host]
	//test: Origin_Uri_No_SubZone()    -> [class:region.zone.host]

}
