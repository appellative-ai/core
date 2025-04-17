package access1

import "time"

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

func Log(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	if originSet {
		defaultLog(&origin, traffic, start, duration, route, req, resp, thresholds)
	} else {
		defaultLog(nil, traffic, start, duration, route, req, resp, thresholds)

	}
}
