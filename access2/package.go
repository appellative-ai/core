package access2

import (
	"github.com/behavioral-ai/core/access"
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

// SetOrigin - initialize the origin
func SetOrigin(o Origin) {
	origin = o
	originSet = true
}

func Log(operators []Operator, traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds access.Threshold) {
	e := NewEvent(traffic, start, duration, route, req, resp, thresholds)
	writeJson(operators, e)
	/*
		if originSet {
			defaultLog(&origin, traffic, start, duration, route, req, resp, thresholds)
		} else {
			defaultLog(nil, traffic, start, duration, route, req, resp, thresholds)

		}

	*/
}
