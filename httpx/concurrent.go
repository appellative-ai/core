package httpx

import (
	"net/http"
	"sync"
	"time"
)

type ConcurrentResult interface {
	Get(name string) *ExchangeResponse
}

type ExchangeInvoke struct {
	Name    string
	Timeout time.Duration
	Do      Exchange
	Req     *http.Request
	Log     func(start time.Time, duration time.Duration, req *http.Request, resp *http.Response)
}

type ExchangeResponse struct {
	Resp *http.Response
	Err  error
}

func DoConcurrent(invokes []ExchangeInvoke) ConcurrentResult {
	var wg sync.WaitGroup

	cnt := len(invokes)
	m := newResults()
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(i *ExchangeInvoke) {
			defer wg.Done()
			var (
				resp   *http.Response
				err    error
				cancel func()
			)
			i.Req, cancel = NewRequestWithTimeout(i.Req, i.Timeout)
			defer cancel()
			resp, err = i.Do(i.Req)
			if err == nil {
				err = TransformBody(resp)
			}
			m.put(i.Name, &ExchangeResponse{Resp: resp, Err: err})
		}(&invokes[i])
	}
	wg.Wait()
	return m
}

type results struct {
	m *sync.Map
}

func newResults() *results {
	c := new(results)
	c.m = new(sync.Map)
	return c
}

func (e *results) Get(name string) *ExchangeResponse {
	value, ok := e.m.Load(name)
	if !ok {
		return nil
	}
	if value1, ok1 := value.(*ExchangeResponse); ok1 {
		return value1
	}
	return nil
}

func (e *results) put(name string, r *ExchangeResponse) {
	e.m.Store(name, r)
}
