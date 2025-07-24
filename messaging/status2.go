package messaging

import "net/http"

var (
	Status2OK = NewStatus2(http.StatusOK, nil)
)

type Status2 interface {
	error
	Code() int
	OK() bool
	Location() string
}

type status2 struct {
	err      error
	code     int
	location string
}

func NewStatus2(code int, err error) *status2 {
	s := new(status2)
	s.code = code
	s.err = err
	return s
}

func (s *status2) Code() int { return s.code }
func (s *status2) OK() bool  { return s.code == http.StatusOK }
func (s *status2) Error() string {
	if s.err != nil {
		return s.err.Error()
	}
	return ""
}
func (s *status2) Location() string { return s.location }
func (s *status2) WithLocation(l string) Status2 {
	s.location = l
	return s
}
