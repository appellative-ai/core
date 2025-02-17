package test

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/iox"
	"io"
	"net/http"
	"testing"
)

func NewRequest(uri any) (*http.Request, *aspect.Status) {
	if uri == nil {
		return nil, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return nil, errors.New(fmt.Sprintf("error: invalid URL scheme : %v", u.Scheme))
	//}
	buf, err := iox.ReadFile(uri)
	if err != nil {
		return nil, aspect.NewStatusError(aspect.StatusIOError, err)
	}
	//status := aspect.StatusOK()
	byteReader := bytes.NewReader(buf)
	reader := bufio.NewReader(byteReader)
	req, err1 := http.ReadRequest(reader)
	if err1 != nil {
		return nil, aspect.NewStatusError(aspect.StatusInvalidArgument, err1)
	}
	bytes1, err2 := ReadContent(buf)
	if err2 != nil {
		return req, aspect.NewStatusError(aspect.StatusIOError, err2)
	}
	if bytes1 != nil {
		req.Body = io.NopCloser(bytes1)
	}
	return req, aspect.StatusOK()
}

func NewRequestTest(uri any, t *testing.T) *http.Request {
	req, status := NewRequest(uri)
	if status.OK() {
		return req
	}
	t.Errorf("ReadRequest() err = %v", status.Err.Error())
	req2, _ := http.NewRequest("", "http://somedomain.com/invalid-uri", nil)
	return req2
}

/*
func createExchange(h http.Header) *aspect.ExchangeOverride {
	if h == nil {
		return nil
	}
	var ex *aspect.ExchangeOverride

	if str, ok := h[httpx.ExchangeOverride]; ok && str[0] != "" {
		for _, s := range str {
			if s == "" {
				continue
			}
			if ex == nil {
				ex = aspect.NewExchangeOverrideEmpty()
			}
			prefix := aspect.ExchangeRequestKey + httpx.ResolverSeparator
			if strings.HasPrefix(s, prefix) {
				ex.SetRequest(s[len(prefix):])
				continue
			}
			prefix = aspect.ExchangeResponseKey + httpx.ResolverSeparator
			if strings.HasPrefix(s, prefix) {
				ex.SetResponse(s[len(prefix):])
				continue
			}
			prefix = aspect.ExchangeStatusKey + httpx.ResolverSeparator
			if strings.HasPrefix(s, prefix) {
				ex.SetStatus(s[len(prefix):])
			}
		}
	}
	return ex
}


*/
