package core

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//https://github.com/advanced-go/core/blob/main/runtime/env.go?line=27

const (
	TimestampName  = "timestamp"
	StatusName     = "status"
	CodeName       = "code"
	TraceName      = "trace"
	RequestIdName  = "request-id"
	ErrorsName     = "errors"
	githubHost     = "github"
	githubDotCom   = "github.com"
	githubTemplate = "https://%v/tree/main%v"
	fragmentId     = "#"
	urnSeparator   = ":"
)

// ErrorHandler - error handler interface
type ErrorHandler interface {
	Handle(s *Status) *Status
}

// LogFunc - log function
type LogFunc func(code int, status, requestId string, errs []error, trace []string)

//type Formatter func(ts time.Time, code int, status, requestId string, errs []error, trace []string) string

var (
	formatter             = defaultFormatter
	logger                = defaultLogger
	defaultLogger LogFunc = func(code int, status, requestId string, errs []error, trace []string) {
		log.Default().Println(formatter(time.Now().UTC(), code, status, requestId, errs, trace))
	}
)

// SetLogFunc - optional override of logging
func SetLogFunc(fn LogFunc) {
	if fn != nil {
		logger = fn
	}
}

// Bypass - bypass error handler
type Bypass struct{}

// Handle - bypass error handler
func (h Bypass) Handle(s *Status) *Status {
	return s
}

// Output - standard output error handler
type Output struct{}

// Handle - output error handler
func (h Output) Handle(s *Status) *Status {
	if s == nil {
		return StatusOK()
	}
	if s.OK() {
		return s
	}
	if s.Err != nil && !s.Handled {
		s.AddParentLocation()
		fmt.Printf("%v", formatter(time.Now().UTC(), s.Code, HttpStatus(s.Code), s.RequestId, []error{s.Err}, s.Trace()))
		s.Handled = true
	}
	return s
}

// Log - log error handler
type Log struct{}

// Handle - log error handler
func (h Log) Handle(s *Status) *Status {
	if s == nil {
		return StatusOK()
	}
	if s.OK() {
		return s
	}
	if s.Err != nil && !s.Handled {
		s.AddParentLocation()
		go logger(s.Code, HttpStatus(s.Code), s.RequestId, []error{s.Err}, s.Trace())
		s.Handled = true
	}
	return s
}

func defaultFormatter(ts time.Time, code int, status, requestId string, errs []error, trace []string) string {
	str := strconv.Itoa(code)
	return fmt.Sprintf("{ %v, %v, %v, %v, %v, %v }\n",
		JsonMarkup(TimestampName, FmtRFC3339Millis(ts), true),
		JsonMarkup(CodeName, str, false),
		JsonMarkup(StatusName, status, true),
		JsonMarkup(RequestIdName, requestId, true),
		formatErrors(ErrorsName, errs),
		formatTrace(TraceName, trace))
}

func formatErrors(name string, errs []error) string {
	if len(errs) == 0 || errs[0] == nil {
		return fmt.Sprintf("\"%v\" : null", name)
	}
	result := fmt.Sprintf("\"%v\" : [ ", name)
	for i, e := range errs {
		if i != 0 {
			result += ","
		}
		result += fmt.Sprintf("\"%v\"", e.Error())
	}
	return result + " ]"
}

func formatTrace(name string, trace []string) string {
	if len(trace) == 0 {
		return fmt.Sprintf("\"%v\" : null", name)
	}
	result := fmt.Sprintf("\"%v\" : [ ", name)
	for i := len(trace) - 1; i >= 0; i-- {
		if i < len(trace)-1 {
			result += ","
		}
		result += fmt.Sprintf("\"%v\"", formatUri(trace[i]))
	}
	return result + " ]"
}

func formatUri(uri string) string {
	i := strings.Index(uri, githubHost)
	if i == -1 {
		return uri
	}
	uri = strings.Replace(uri, githubHost, githubDotCom, len(githubDotCom))
	i = strings.LastIndex(uri, "/")
	if i != -1 {
		first := uri[:i]
		last := uri[i:]
		last = strings.Replace(last, urnSeparator, fragmentId, len(fragmentId))
		return fmt.Sprintf(githubTemplate, first, last)
	}
	return uri
}

// NewInvalidBodyTypeError - invalid type error
func NewInvalidBodyTypeError(t any) error {
	return errors.New(fmt.Sprintf("invalid body type: %v", reflect.TypeOf(t)))
}
