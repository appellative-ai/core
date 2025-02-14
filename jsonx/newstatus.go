package jsonx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"strings"
)

const (
	StatusOKUri       = "urn:status:ok"
	StatusNotFoundUri = "urn:status:notfound"
	StatusTimeoutUri  = "urn:status:timeout"
	statusToken       = "status"
)

type serializedStatusState struct {
	Code     int    `json:"code"`
	Location string `json:"location"`
	Err      string `json:"err"`
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
func NewStatusFrom(uri string) error {
	status := statusFromConst(uri)
	if status != nil {
		return status
	}
	//status = ValidateUri(uri)
	//if !status.OK() {
	//	return status
	//}
	buf, status1 := iox.ReadFile(uri) //iox.FileName(uri))
	if status1 != nil {
		return status1 //aspect.NewStatusError(aspect.StatusIOError, err1)
	}
	var status2 serializedStatusState
	err := json.Unmarshal(buf, &status2)
	if err != nil {
		return err //aspect.NewStatusError(aspect.StatusJsonDecodeError, err)
	}
	if len(status2.Err) > 0 {
		return errors.New(status2.Err) //aspect.NewStatusError(status2.Code, errors.New(status2.Err))
	}
	return errors.New(fmt.Sprintf("code : %v", status2.Code)) //aspect.NewStatus(status2.Code).AddLocation()
}

func statusFromConst(url string) error {
	if len(url) == 0 {
		return nil
	}
	switch url {
	case StatusOKUri:
		return nil //aspect.StatusOK()
	case StatusNotFoundUri:
		return errors.New("not found") //aspect.NewStatus(http.StatusNotFound)
	case StatusTimeoutUri:
		return errors.New("gatewar timeout") //aspect.NewStatus(http.StatusGatewayTimeout)
	}
	return nil
}
