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
	origin    = Origin{}
	originSet bool
)

func init() {
	log.SetFlags(0)
}

// SetOrigin - initialize the origin
func SetOrigin(o Origin) {
	origin = o
	originSet = true
}

func Log(operators []Operator, traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	e := newEvent(traffic, start, duration, route, req, resp, thresholds)
	if len(operators) == 0 {
		operators = defaultOperators
	}
	s := writeJson(operators, e)
	log.Printf("%v\n", s)
}
