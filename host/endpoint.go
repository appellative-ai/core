package host

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"net/http"
)

func ExchangeHandler(w http.ResponseWriter, req *http.Request, resp *http.Response) {
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, req.Header)
}

func Init(r *http.Request) {
	httpx.AddRequestId(r)
}

func NewEndpoint(links ...any) *rest.Endpoint {
	chain := rest.BuildChain(links...)
	return rest.NewEndpoint(ExchangeHandler, Init, chain)
}
