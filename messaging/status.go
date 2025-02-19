package messaging

import (
	"fmt"
	"net/http"
)

const (
	StatusIOError         = int(91) // I/O operation failed
	StatusJsonDecodeError = int(92) // Json decoding failed
	StatusJsonEncodeError = int(93) // Json encoding failed
)

type Status struct {
	Code int   `json:"code"`
	Err  error `json:"err"`
}

func StatusOK() *Status {
	return okStatus
}

func StatusBadRequest() *Status {
	return badRequestStatus
}

func StatusNotFound() *Status {
	return notFoundStatus
}

func NewStatus(code int) *Status {
	s := new(Status)
	s.Code = code
	return s
}

func NewStatusError(code int, err error) *Status {
	s := new(Status)
	s.Code = code
	s.Err = err
	return s
}

func (s *Status) OK() bool {
	return s.Code == http.StatusOK
}

func (s *Status) NotFound() bool {
	return s.Code == http.StatusNotFound
}

func (s *Status) String() string {
	if s.Err != nil {
		return fmt.Sprintf("%v [%v]", s.Code, s.Err)
	} else {
		return fmt.Sprintf("%v", s.Code)
	}
}

var okStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusOK
	return s
}()

var badRequestStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusBadRequest
	return s
}()

var notFoundStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusNotFound
	return s
}()
