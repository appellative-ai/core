package access2

import (
	"log"
	"time"
)

// InternalTraffic = "internal"

const (
	EgressTraffic  = "egress"
	IngressTraffic = "ingress"

	failsafeUri     = "https://invalid-uri.com"
	XRequestId      = "x-request-id"
	XRateLimit      = "x-rate-limit"
	XRateBurst      = "x-rate-burst"
	XTimeout        = "x-timeout"
	XRedirect       = "x-redirect"
	XCached         = "x-cached"
	ContentEncoding = "Content-Encoding"
)

var (
	origin           = originT{}
	originSet        bool
	defaultOperators []Operator
)

func init() {
	log.SetFlags(0)
}

// SetOrigin -
func SetOrigin(region, zone, subZone, host, instanceId string) {
	origin.Region = region
	origin.Zone = zone
	origin.SubZone = subZone
	origin.Host = host
	origin.InstanceId = instanceId
	originSet = true
}

// SetOperators - initialize the operators
func SetOperators(o []Operator) {
	if len(o) > 0 {
		defaultOperators = o
	}
}

func Log(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	LogWithOperators(defaultOperators, traffic, start, duration, route, req, resp, thresholds)
}

func LogWithOperators(operators []Operator, traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	if len(operators) == 0 {
		log.Printf("%v\n", "{ \"error\" : \"no operators configured\" }")
		return
	}
	e := newEvent(traffic, start, duration, route, req, resp, thresholds)
	s := writeJson(operators, e)
	log.Printf("%v\n", s)
}
