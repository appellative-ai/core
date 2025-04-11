package host

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"net/http"
)

func Exchange(w http.ResponseWriter, r *http.Request, handler rest.Exchange) {
	httpx.AddRequestId(r)
	if handler == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, _ := handler(r)
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}

func ExchangeHandler(w http.ResponseWriter, req *http.Request, resp *http.Response) {
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, req.Header)
}
