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

/*
resp, err = authExchange(r)
	if !okFunc(resp.StatusCode) {
		w.WriteHeader(resp.StatusCode)
		access.Log(access.IngressTraffic, start, time.Since(start), r, resp, access.Controller{Timeout: hostDuration, RateLimit: "0", RateBurst: "0", Code: controllerCode})
		return
	}
	from := r.Header.Get(XFrom)
	if from == "" {
		r.Header.Set(XFrom, Route)
	}
	r.Header.Set(XFrom, Route)
	// TODO: Need to create a new request with the appropriate timeout and host name, using an intermediary
	if hostDuration > 0 {
		ctx, cancel := context.WithTimeout(r.Context(), hostDuration)
		defer cancel()
		r2 := r.Clone(ctx)
		resp, err = handler(r2)
	} else {
		resp, err = handler(r)
	}
	resp.Header.Del(XRoute)
	if err != nil && err.Error() == "http.StatusGatewayTimeout" {
		controllerCode = access.ControllerTimeout
	}
	resp.ContentLength = httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
	r.Header.Set(XTo, Route)
	access.Log(access.IngressTraffic, start, time.Since(start), r, resp, access.Controller{Timeout: hostDuration, RateLimit: "0", RateBurst: "0", Code: controllerCode})
}

*/
