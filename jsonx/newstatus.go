package jsonx

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/iox"
	"net/http"
	"strings"
)

const (
	StatusOKUri       = "urn:status:ok"
	StatusNotFoundUri = "urn:status:notfound"
	StatusTimeoutUri  = "urn:status:timeout"
	statusToken       = "status"
)

type serializedStatusState struct {
	Code     int    `jsonx:"code"`
	Location string `jsonx:"location"`
	Err      string `jsonx:"err"`
}

// isStatusURL - determine if the file name of the URL contains the text 'status'
func isStatusURL(url string) bool {
	if len(url) == 0 {
		return false
	}
	i := strings.LastIndex(url, statusToken)
	if i == -1 {
		return false
	}
	return strings.LastIndex(url, "/") < i
}

// NewStatusFrom - create a new Status from a URI
func NewStatusFrom(uri string) *aspect.Status {
	status := statusFromConst(uri)
	if status != nil {
		return status
	}
	//status = ValidateUri(uri)
	//if !status.OK() {
	//	return status
	//}
	buf, status1 := iox.ReadFile(uri) //iox.FileName(uri))
	if !status1.OK() {
		return status1 //aspect.NewStatusError(aspect.StatusIOError, err1)
	}
	var status2 serializedStatusState
	err := json.Unmarshal(buf, &status2)
	if err != nil {
		return aspect.NewStatusError(aspect.StatusJsonDecodeError, err)
	}
	if len(status2.Err) > 0 {
		return aspect.NewStatusError(status2.Code, errors.New(status2.Err))
	}
	return aspect.NewStatus(status2.Code).AddLocation()
}

func statusFromConst(url string) *aspect.Status {
	if len(url) == 0 {
		return aspect.StatusOK()
	}
	switch url {
	case StatusOKUri:
		return aspect.StatusOK()
	case StatusNotFoundUri:
		return aspect.NewStatus(http.StatusNotFound)
	case StatusTimeoutUri:
		return aspect.NewStatus(http.StatusGatewayTimeout)
	}
	return nil
}
