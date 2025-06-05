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

// OriginT - location
type OriginT struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
	Collective string
	Domain     string
}

func (o OriginT) String() string { return o.Name() }

func (o OriginT) Name() string {
	var name = fmt.Sprintf(originNameFmt, o.Collective, o.Domain, ServiceKind)

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

func NewOriginFromMessage(m *Message, collective, domain string) (OriginT, *Status) {
	cfg, status := MapContent(m)
	if !status.OK() {
		return OriginT{}, status
	}
	return NewOrigin(cfg, collective, domain)
}

func NewOrigin(m map[string]string, collective, domain string) (OriginT, *Status) {
	var origin OriginT
	if domain == "" || collective == "" {
		return origin, NewStatus(StatusInvalidArgument, errors.New("error: origin collective or domain is empty"))
	}
	origin.Domain = domain
	origin.Collective = collective

	if m == nil {
		return origin, NewStatus(StatusInvalidArgument, errors.New("error: origin map is nil"))
	}
	origin.Region = m[RegionKey]
	if origin.Region == "" {
		return origin, NewStatus(StatusInvalidContent, errors.New(fmt.Sprintf("config map does not contain key: %v", RegionKey)))
	}

	origin.Zone = m[ZoneKey]
	if origin.Zone == "" {
		return origin, NewStatus(StatusInvalidContent, errors.New(fmt.Sprintf("config map does not contain key: %v", ZoneKey)))
	}

	origin.SubZone = m[SubZoneKey]

	origin.Host = m[HostKey]
	if origin.Host == "" {
		return origin, NewStatus(StatusInvalidContent, errors.New(fmt.Sprintf("config map does not contain key: %v", HostKey)))
	}
	origin.InstanceId = m[InstanceIdKey]

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
