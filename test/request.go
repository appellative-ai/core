package test

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/iox"
	"io"
	"net/http"
	"testing"
)

func NewRequest(uri any) (*http.Request, *core.Status) {
	if uri == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return nil, errors.New(fmt.Sprintf("error: invalid URL scheme : %v", u.Scheme))
	//}
	buf, status := iox.ReadFile(uri)
	if !status.OK() {
		return nil, status
	}
	byteReader := bytes.NewReader(buf)
	reader := bufio.NewReader(byteReader)
	req, err1 := http.ReadRequest(reader)
	if err1 != nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, err1)
	}
	bytes1, err2 := ReadContent(buf)
	if err2 != nil {
		return req, core.NewStatusError(core.StatusIOError, err2)
	}
	if bytes1 != nil {
		req.Body = io.NopCloser(bytes1)
	}
	/*
		ex := createExchange(req.Header)
		if ex != nil {
			ctx := core.NewExchangeOverrideContext(nil, ex)
			req2, err3 := http.NewRequestWithContext(ctx, req.Method, req.URL.String(), req.Body)
			if err3 != nil {
				return nil, core.NewStatusError(core.StatusInvalidArgument, err3)
			}
			req2.Header = req.Header
			return req2, core.StatusOK()
		}

	*/
	return req, core.StatusOK()
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
func createExchange(h http.Header) *core.ExchangeOverride {
	if h == nil {
		return nil
	}
	var ex *core.ExchangeOverride

	if str, ok := h[httpx.ExchangeOverride]; ok && str[0] != "" {
		for _, s := range str {
			if s == "" {
				continue
			}
			if ex == nil {
				ex = core.NewExchangeOverrideEmpty()
			}
			prefix := core.ExchangeRequestKey + httpx.ResolverSeparator
			if strings.HasPrefix(s, prefix) {
				ex.SetRequest(s[len(prefix):])
				continue
			}
			prefix = core.ExchangeResponseKey + httpx.ResolverSeparator
			if strings.HasPrefix(s, prefix) {
				ex.SetResponse(s[len(prefix):])
				continue
			}
			prefix = core.ExchangeStatusKey + httpx.ResolverSeparator
			if strings.HasPrefix(s, prefix) {
				ex.SetStatus(s[len(prefix):])
			}
		}
	}
	return ex
}


*/
