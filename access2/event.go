package access2

import (
	"fmt"
	"github.com/behavioral-ai/core/access"
	"github.com/behavioral-ai/core/fmtx"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Event - struct for all access logging
type Event struct {
	Traffic    string
	Start      time.Time
	Duration   time.Duration
	Route      string
	Req        any
	Resp       any
	Thresholds access.Threshold
	NewReq     *http.Request
	NewResp    *http.Response
	Url        string
	Parsed     *access.Parsed
}

func NewEvent(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds access.Threshold) *Event {
	e := new(Event)
	e.Traffic = traffic
	e.Start = start
	e.Duration = duration
	e.Route = route
	e.Req = req
	e.Resp = resp
	e.Thresholds = thresholds
	e.NewReq = access.BuildRequest(req)
	e.NewResp = access.BuildResponse(resp)
	e.Url, e.Parsed = access.ParseURL(e.NewReq.Host, e.NewReq.URL)
	return e
}

func (e *Event) Value(value string) string {
	switch value {
	case TrafficOperator:
		return e.Traffic
	case StartTimeOperator:
		return fmtx.FmtRFC3339Millis(e.Start) //strings2.FmtTimestamp(l.Start)
	case DurationOperator:
		//d := int(l.Duration / time.Duration(1e6))
		//return strconv.Itoa(d)
		return strconv.Itoa(fmtx.Milliseconds(e.Duration))
	case DurationStringOperator:
		return "" //l.Duration.String()
	case RouteOperator:
		return e.Route

		// Origin
	case OriginRegionOperator:
		return "" //runtime.OriginRegion()
	case OriginZoneOperator:
		return "" //runtime.OriginZone()
	case OriginSubZoneOperator:
		return "" //runtime.OriginSubZone()
	case OriginServiceOperator:
		return "" //runtime.OriginService()
	case OriginInstanceIdOperator:
		return "" //runtime.OriginInstanceId()

		// Request
	case RequestMethodOperator:
		return e.NewReq.Method
	case RequestProtocolOperator:
		return e.NewReq.Proto
	case RequestPathOperator:
		return e.NewReq.URL.Path
	case RequestUrlOperator:
		return e.NewReq.URL.String()
	case RequestHostOperator:
		return e.NewReq.Host
	case RequestIdOperator:
		return "" //l.Header.Get(RequestIdHeaderName)
	case RequestFromRouteOperator:
		return "" //l.Header.Get(FromRouteHeaderName)
	case RequestUserAgentOperator:
		return e.NewReq.Header.Get(UserAgentHeaderName)
	case RequestAuthorityOperator:
		return ""
	case RequestForwardedForOperator:
		return e.NewReq.Header.Get(ForwardedForHeaderName)

		// Response
	case StatusFlagsOperator:
		return "" //l.StatusFlags
	case ResponseBytesReceivedOperator:
		return strconv.Itoa(int(0)) //l.BytesReceived))
	case ResponseBytesSentOperator:
		return fmt.Sprintf("%v", 0) //l.BytesSent)
	case ResponseStatusCodeOperator:
		return strconv.Itoa(e.NewResp.StatusCode)

	// Controller State
	case ControllerNameOperator:
		return "" //l.ControllerName
	case TimeoutDurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.Thresholds.TimeoutT())) //strconv.Itoa(l.Timeout)
	case RateLimitOperator:
		return fmt.Sprintf("%v", e.Thresholds.RateLimitT())
	case RateBurstOperator:
		return "" //strconv.Itoa(l.RateBurst)
	case RedirectOperator:
		return strconv.Itoa(e.Thresholds.RedirectT())
		//case ProxyThresholdOperator:
		//	return l.ProxyThreshold
		//case RetryOperator:
		//	return l.Retry
		//case RetryRateLimitOperator:
		//		return l.CtrlState[RetryRateLimitName]
		//	case RetryRateBurstOperator:
		//		return l.CtrlState[RetryRateBurstName]
	}
	if strings.HasPrefix(value, RequestReferencePrefix) {
		name := requestOperatorHeaderName(value)
		return e.NewReq.Header.Get(name)
	}
	if !strings.HasPrefix(value, OperatorPrefix) {
		return value
	}
	return ""
}
