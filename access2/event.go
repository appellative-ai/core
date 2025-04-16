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
	Thresholds Threshold
	NewReq     *http.Request
	NewResp    *http.Response
	Url        string
	Parsed     *access.Parsed
}

func NewEvent(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) *Event {
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

func (e *Event) AddRequest(r *http.Request) {
	e.NewReq = access.BuildRequest(r)
}

func (e *Event) AddResponse(r *http.Response) {
	e.NewResp = access.BuildResponse(r)
}

func (e *Event) Value(value string) string {
	switch value {
	case TrafficOperator:
		return e.Traffic
	case StartTimeOperator:
		return fmtx.FmtRFC3339Millis(e.Start)
	case DurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.Duration))
	case DurationStringOperator:
		return e.Duration.String()
	case RouteOperator:
		return e.Route

		// Origin
	case OriginRegionOperator:
		return origin.Region
	case OriginZoneOperator:
		return origin.Zone
	case OriginSubZoneOperator:
		return origin.SubZone
	case OriginServiceOperator:
		return origin.Host
	case OriginInstanceIdOperator:
		return origin.InstanceId

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
		return e.NewReq.Header.Get(RequestIdHeaderName)
	case RequestFromRouteOperator:
		return e.NewReq.Header.Get(FromRouteHeaderName)
	case RequestUserAgentOperator:
		return e.NewReq.Header.Get(UserAgentHeaderName)
	case RequestAuthorityOperator:
		return ""
	case RequestForwardedForOperator:
		return e.NewReq.Header.Get(ForwardedForHeaderName)

		// Response
	case ResponseBytesReceivedOperator:
		return fmt.Sprintf("%v", e.NewResp.ContentLength) //strconv.Itoa(e.NewResp.ContentLength) //l.BytesReceived))
	case ResponseBytesSentOperator:
		return fmt.Sprintf("%v", 0) //l.BytesSent)
	case ResponseStatusCodeOperator:
		if e.NewResp == nil {
			return "0"
		} else {
			return strconv.Itoa(e.NewResp.StatusCode)
		}
	case ResponseContentEncodingOperator:
		return access.Encoding(e.NewResp)
	case ResponseCachedOperator:
		s := e.NewResp.Header.Get(XCached)
		if s == "" {
			s = "false"
		}
		return s

	// Thresholds
	case TimeoutDurationOperator:
		return strconv.Itoa(fmtx.Milliseconds(e.Thresholds.TimeoutT())) //strconv.Itoa(l.Timeout)
	case RateLimitOperator:
		return fmt.Sprintf("%v", e.Thresholds.RateLimitT())
	case RedirectOperator:
		return strconv.Itoa(e.Thresholds.RedirectT())
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
