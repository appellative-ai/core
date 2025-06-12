package host

import (
	"fmt"
	access "github.com/behavioral-ai/core/access2"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"net/http"
	"net/http/httptest"
	"time"
)

type agentT struct{}

func newTestAgent() *agentT {
	return new(agentT)
}
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return "agent:test" }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {}

func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (*http.Response, error) {
		return nil, nil
	}
}

func ExchangeTest(w http.ResponseWriter, r *http.Request, handler rest.Exchange) {
	httpx.AddRequestId(r)
	if handler == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, _ := handler(r)
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}

func _ExampleHost() {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8081/github/advanced-go/search:google?q=golang", nil)

	ExchangeTest(r, req, func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK}, nil
	})

	fmt.Printf("test: Exchange() -> [resp:%v]\n", r.Result().StatusCode)

	//Output:
	//test: Exchange() -> [resp:200]

}

const (
	authorization = "Authorization"
	route         = "host"
)

func authorizationLink(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(authorization)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		resp, err = next(r)
		return
	}
}

func accessLogLink(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		limit := ""
		pct := ""
		timeout := ""

		if next != nil {
			resp, err = next(r)
		}
		limit = resp.Header.Get(access.XRateLimit)
		resp.Header.Del(access.XRateLimit)
		timeout = resp.Header.Get(access.XTimeout)
		resp.Header.Del(access.XTimeout)
		pct = resp.Header.Get(access.XRedirect)
		resp.Header.Del(access.XRedirect)
		access.Agent.Log(access.IngressTraffic, start, time.Since(start), route, r, resp, access.Threshold{Timeout: timeout, RateLimit: limit, Redirect: pct})
		return
	}
}

func ExampleNewEndpoint() {
	agent := newTestAgent()
	fmt.Printf("test: NewEndpoint() -> [%v]\n", agent)

	e := NewEndpoint("/resource", []any{accessLogLink, authorizationLink, agent})
	fmt.Printf("test: NewEndpoint() -> [%v]\n", e)

	//Output:
	//test: NewEndpoint() -> [agent:test]
	//test: NewEndpoint() -> [&{/resource 0xfa0000 0xfa1d20 0xfa0120}]

}

/*
func ExampleExchangeHandler() {
	e := NewEndpoint2(nil)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	e.Exchange(rec, req)

	fmt.Printf("test: ExchangeHandler() -> [%v]\n", req.URL.String())

	//Output:
	//test: ExchangeHandler() -> [https://www.google.com/search?q=golang]

}


*/
