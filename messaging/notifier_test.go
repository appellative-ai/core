package messaging

import (
	"errors"
	"net/http"
)

func ExampleNotify() {
	s := NewStatusError(http.StatusGatewayTimeout, errors.New("rate limiting"), "message", "test:agent")
	s.AgentUri = "resiliency:agent/operative"

	Notify(s)

	//Output:
	//notify-> 2025-02-26T15:34:45.784Z [resiliency:agent/operative] [core:messaging.status] [Timeout] [rate limiting]

}
