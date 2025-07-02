package rest

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func UpdateExchange(name string, ex *func(r *http.Request) (*http.Response, error), m *messaging.Message) {
	if m == nil || ex == nil || m.ContentType() != ContentTypeExchange {
		return
	}
	newEx, status := ExchangeContent(m)
	if !status.OK() {
		messaging.Reply(m, status, name)
		return
	}
	*ex = newEx
}
