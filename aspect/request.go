package aspect

import "net/http"

const (
	XRequestId = "X-Request-Id"
)

// RequestId - return a request id from any type and will create a new one if not found
func RequestId(t any) string {
	if t == nil {
		//s, _ := uuid.NewUUID()
		return "" // s.String()
	}
	requestId := ""
	switch ptr := t.(type) {
	case string:
		requestId = ptr
	case *http.Request:
		requestId = ptr.Header.Get(XRequestId)
	case http.Header:
		requestId = ptr.Get(XRequestId)
	}
	//if len(requestId) == 0 {
	//	s, _ := uuid.NewUUID()
	//	requestId = s.String()
	//}
	return requestId
}
