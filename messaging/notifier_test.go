package messaging

import (
	"errors"
	"net/http"
)

func ExampleNotify() {
	s := NewStatusError(http.StatusGatewayTimeout, errors.New("rate limiting"), "message", nil)
	s.AgentUri = "resiliency:agent/operative"

	Notify(s)

	//Output:
	//notify-> [] [core:messaging.status] [] [Not Found - error: not found]

}
