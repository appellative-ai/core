package httpx

import (
	"errors"
	"fmt"
	//"github.com/advanced-go/stdlib/core"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"sync"
	"sync/atomic"
)

type ExchangeResultM struct {
	Failure bool
	Resp    *http.Response
	Status  error //*messaging.Status
}

type OnResponse func(resp *http.Response, status *messaging.Status) (failure, proceed bool)

func MultiExchange(reqs []*http.Request, handler OnResponse) ([]ExchangeResultM, *messaging.Status) {
	cnt := len(reqs)
	if cnt == 0 || handler == nil {
		fmt.Printf("%v", "error: no requests were found to process, or OnResponse handler is nil")
		return nil, messaging.NewStatusError(messaging.StatusInvalidArgument, errors.New("error: no requests were found to process"), "")
	}
	var wg sync.WaitGroup
	failure := atomic.Bool{}

	results := make([]ExchangeResultM, cnt)
	for i := 0; i < cnt; i++ {
		if reqs[i] == nil {
			continue
		}
		wg.Add(1)
		go func(req *http.Request, res *ExchangeResultM) {
			defer wg.Done()
			//res.Resp, res.Status = Exchange(req)
			if handler != nil {
				fail, proceed := handler(res.Resp, nil)
				if fail {
					res.Failure = true
					failure.Store(true)
				}
				if !proceed {
					return
				}
			}
		}(reqs[i], &results[i])
	}
	wg.Wait()
	if failure.Load() {
		return results, messaging.NewStatusError(messaging.StatusExecError, errors.New("error: request failures"), "")
	}
	return results, messaging.StatusOK()
}
