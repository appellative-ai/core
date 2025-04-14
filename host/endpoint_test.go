package host

import (
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
	"net/http"
	"net/http/httptest"
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

func (a *agentT) Exchange(r *http.Request) (*http.Response, error) {
	return nil, nil
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

func ExampleNewEndpoint() {
	agent := newTestAgent()
	fmt.Printf("test: NewEndpoint() -> [%v]\n", agent)

	rest.BuildChain(AccessLogLink, AuthorizationLink, agent)
	e := NewEndpoint(agent)
	fmt.Printf("test: NewEndpoint() -> [%v]\n", e)

	//Output:
	//test: NewEndpoint() -> [agent:test]
	//test: NewEndpoint() -> [&{0xa2ee00 0xa2f460 0xa2ef20}]

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
