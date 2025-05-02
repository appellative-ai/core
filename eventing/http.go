package eventing

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func httpAddActivity(hostName, uri, event, source string, content any) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}

func httpNotify(e NotifyEvent) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}
