package messaging

import (
	"fmt"
	"net/http"
)

var (
	Status2OK       = NewStatus2(http.StatusOK, "", nil)
	Status2NotFound = NewStatus2(http.StatusNotFound, "", nil)
)

/*
type Status2 interface {
	//fmt.Stringer
	//error
	String() string
	Code() int
	OK() bool
	Location() string
}


*/

type Status2 struct {
	Err      error
	Code     int
	Location string
}

func NewStatus2(code int, location string, err error) *Status2 {
	s := new(Status2)
	s.Code = code
	s.Err = err
	s.Location = location
	return s
}

func (s *Status2) OK() bool       { return s.Code == http.StatusOK }
func (s *Status2) NotFound() bool { return s.Code == http.StatusNotFound }
func (s *Status2) String() string {
	if s.Err != nil {
		return fmt.Sprintf("%v - %v", HttpStatus(s.Code), s.Err)
	}
	return fmt.Sprintf("%v", HttpStatus(s.Code))

}

// func (s *Status2) Code() int { return s.Code }
// func (s *Status2) Error() error { return s.Err }
// func (s *status2) Location() string { return s.location }
