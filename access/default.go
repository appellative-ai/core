package access

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
	//fmt2 "github.com/advanced-go/stdlib/fmt"
	"github.com/behavioral-ai/core/uri"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var defaultLog = func(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) {
	s := DefaultFormat(o, traffic, start, duration, req, resp, routing, controller)
	log.Default().Printf("%v\n", s)
}

func DefaultFormat(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) string {
	newReq := BuildRequest(req)
	newResp := BuildResponse(resp)
	url, parsed := uri.ParseURL(newReq.Host, newReq.URL)
	o.Host = Conditional(o.Host, parsed.Host)
	if controller.RateLimit == 0 {
		controller.RateLimit = -1
	}
	if controller.RateBurst == 0 {
		controller.RateBurst = -1
	}
	s := fmt.Sprintf("{"+
		"\"region\":%v, "+
		"\"zone\":%v, "+
		"\"sub-zone\":%v, "+
		"\"instance-id\":%v, "+
		"\"traffic\":\"%v\", "+
		"\"start\":%v, "+
		"\"duration\":%v, "+
		"\"request-id\":%v, "+
		"\"relates-to\":%v, "+
		"\"location\":%v, "+
		"\"protocol\":%v, "+
		"\"method\":%v, "+
		"\"host\":%v, "+
		"\"from\":%v, "+
		"\"to\":%v, "+
		"\"uri\":%v, "+
		"\"path\":%v, "+
		"\"query\":%v, "+
		"\"status-code\":%v, "+
		"\"encoding\":%v, "+
		"\"bytes\":%v, "+
		"\"timeout\":%v, "+
		"\"rate-limit\":%v, "+
		"\"rate-burst\":%v, "+
		"\"cc\":%v, "+
		"\"route\":%v, "+
		"\"route-to\":%v, "+
		"\"route-percent\":%v, "+
		"\"rc\":%v }",

		// Origin, traffic, timestamp, duration
		JsonString(o.Region),
		JsonString(o.Zone),
		JsonString(o.SubZone),
		JsonString(o.InstanceId),
		traffic,
		core.FmtRFC3339Millis(start),
		strconv.Itoa(Milliseconds(duration)),

		// Request
		JsonString(newReq.Header.Get(XRequestId)),
		JsonString(newReq.Header.Get(XRelatesTo)),
		JsonString(newReq.Header.Get(LocationHeader)),
		JsonString(newReq.Proto),
		JsonString(newReq.Method),
		JsonString(o.Host),
		JsonString(routing.From),
		JsonString(CreateTo(newReq)),
		JsonString(url),
		JsonString(parsed.Path),
		JsonString(parsed.Query),

		// Response
		newResp.StatusCode,
		//jsonString(resp.Status),
		JsonString(Encoding(newResp)),
		fmt.Sprintf("%v", newResp.ContentLength),

		// Controller
		Milliseconds(controller.Timeout),
		fmt.Sprintf("%v", controller.RateLimit),
		strconv.Itoa(controller.RateBurst),
		JsonString(controller.Code),

		// Routing
		JsonString(routing.Route),
		JsonString(routing.To),
		fmt.Sprintf("%v", routing.Percent),
		JsonString(routing.Code),
	)

	return s
}

// Milliseconds - convert time.Duration to milliseconds
func Milliseconds(duration time.Duration) int {
	if duration <= 0 {
		return -1
	}
	return int(duration / time.Duration(1e6))
}

func BuildRequest(r any) *http.Request {
	if r == nil {
		newReq, _ := http.NewRequest("", failsafeUri, nil)
		return newReq
	}
	if req, ok := r.(*http.Request); ok {
		return req
	}
	if req, ok := r.(Request); ok {
		newReq, _ := http.NewRequest(req.Method(), req.Url(), nil)
		newReq.Header = req.Header()
		newReq.Proto = req.Protocol()
		return newReq
	}
	newReq, _ := http.NewRequest("", "https://somehost.com/search?q=test", nil)
	return newReq
}

func BuildResponse(r any) *http.Response {
	if r == nil {
		newResp := &http.Response{StatusCode: http.StatusOK}
		return newResp
	}
	if newResp, ok := r.(*http.Response); ok {
		return newResp
	}
	if sc, ok := r.(int); ok {
		return &http.Response{StatusCode: sc}
	}
	if status, ok := r.(*core.Status); ok {
		return &http.Response{StatusCode: status.HttpCode()}
	}
	newResp := &http.Response{StatusCode: http.StatusOK}
	return newResp
}

func Encoding(resp *http.Response) string {
	encoding := ""
	if resp != nil && resp.Header != nil {
		encoding = resp.Header.Get(ContentEncoding)
	}
	// normalize encoding
	if strings.Contains(strings.ToLower(encoding), "none") {
		encoding = ""
	}
	return encoding
}

func Conditional(primary, secondary string) string {
	if len(primary) == 0 {
		return secondary
	}
	return primary
}

/*
func threshold(threshold any) int {
	if threshold == nil {
		return 0
	}
	if dur, ok := threshold.(time.Duration); ok {
		return Milliseconds(dur)
	}
	if i, ok1 := threshold.(int); ok1 {
		return i
	}
	if f, ok2 := threshold.(float64); ok2 {
		return int(f)
	}
	if ctx, ok := threshold.(context.Context); ok {
		if deadline, ok1 := ctx.Deadline(); ok1 {
			return Milliseconds(time.Until(deadline))
		}
	}
	return 0
}


*/

func CreateTo(req *http.Request) string {
	if req == nil {
		return ""
	}
	to := req.Header.Get(core.XTo)
	if to != "" {
		return to
	}
	return uri.UprootDomain(req.URL)
}
