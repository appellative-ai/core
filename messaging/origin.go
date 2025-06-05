package messaging

import (
	"errors"
	"fmt"
)

const (
	originNameFmt = "%v:%v:%v"
	ServiceKind   = "service"
	RegionKey     = "region"
	ZoneKey       = "zone"
	SubZoneKey    = "sub-zone"
	HostKey       = "host"
	InstanceIdKey = "instance-id"
)

// Origin - location
type Origin struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func (o Origin) String() string {
	return o.Name("any", "any")
}

func (o Origin) Name(collective, domain string) string {
	var name = fmt.Sprintf(originNameFmt, collective, domain, ServiceKind)

	if o.Region != "" {
		name += "/" + o.Region
	}
	if o.Zone != "" {
		name += "/" + o.Zone
	}
	if o.SubZone != "" {
		name += "/" + o.SubZone
	}
	if o.Host != "" {
		name += "/" + o.Host
	}
	if o.InstanceId != "" {
		name += "#" + o.InstanceId
	}
	return name
}

func NewOriginFromMessage(m *Message) (Origin, *Status) {
	var origin Origin

	cfg, status := MapContent(m)
	if cfg == nil {
		return origin, status
	}

	origin.Region = cfg[RegionKey]
	if origin.Region == "" {
		return origin, NewStatus(StatusInvalidContent, errors.New(fmt.Sprintf("config map does not contain key: %v", RegionKey)))
	}

	origin.Zone = cfg[ZoneKey]
	if origin.Zone == "" {
		return origin, NewStatus(StatusInvalidContent, errors.New(fmt.Sprintf("config map does not contain key: %v", ZoneKey)))
	}

	origin.SubZone = cfg[SubZoneKey]

	origin.Host = cfg[HostKey]
	if origin.Host == "" {
		return origin, NewStatus(StatusInvalidContent, errors.New(fmt.Sprintf("config map does not contain key: %v", HostKey)))
	}
	origin.InstanceId = cfg[InstanceIdKey]

	return origin, nil
}

/*
//if o.SubZone == "" {
	//	messaging.Reply(m, messaging.ConfigMapContentError(nil, SubZoneKey), NamespaceName)
	//	return
	//}

	if o.InstanceId == "" {
		messaging.Reply(m, messaging.ConfigMapContentError(a, InstanceIdKey), a.Name())
		return
	}



*/

/*
func (o Origin) Uri(class string) string {
	return fmt.Sprintf(uriFmt, class, o)
}

func (o Origin) String() string { return "" }


*/

/*

var uri = o.Region

if o.Zone != "" {
uri += "." + o.Zone
}
if o.SubZone != "" {
uri += "." + o.SubZone
}
if o.Host != "" {
uri += "." + o.Host
}
return uri
}


*/
