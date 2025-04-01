package host

import (
	"github.com/behavioral-ai/core/httpx"
	"net/http"
)

func Exchange(w http.ResponseWriter, r *http.Request, handler httpx.Exchange) {
	httpx.AddRequestId(r)
	if handler == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, _ := handler(r)
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}
