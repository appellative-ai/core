package rest

import (
	"github.com/behavioral-ai/core/messaging"
)

func UpdateExchange(name string, ex *Exchange, m *messaging.Message) {
	if m == nil || ex == nil || m.ContentType() != ContentTypeExchange {
		return
	}
	newEx, status := ExchangeContent(m)
	if !status.OK() {
		messaging.Reply(m, status, name)
		return
	}
	*ex = newEx
	messaging.Reply(m, messaging.StatusOK(), name)
}

/*
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


*/
