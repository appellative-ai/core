package http

import (
	"errors"
	"fmt"
	"net/http"
)

type FramedExchange func(req *http.Request, next *Frame) (*http.Response, error)

type Frame struct {
	Name string
	Fn   FramedExchange
	Next *Frame
}

func (f *Frame) String() string {
	return f.Name
}

type ExchangePipeline struct {
	head *Frame
}

func NewExchangePipeline(fn ...FramedExchange) *ExchangePipeline {
	e := new(ExchangePipeline)
	if len(fn) == 0 {
		return e
	}
	var prev *Frame

	i := 0
	for i = len(fn) - 1; i >= 0; i-- {
		f := new(Frame)
		f.Name = fmt.Sprintf("do%v", i+1)
		f.Fn = fn[i]
		if prev == nil {
			prev = f
		} else {
			f.Next = prev
			prev = f
		}
		//fmt.Printf("test: NewExchangePipeline() -> [curr:%v] [next:%v]\n", f, f.Next)
		if e.head == nil {
			e.head = f
		} else {
			t := e.head
			e.head = f
			e.head.Next = t
		}
	}
	return e
}

func (e *ExchangePipeline) Run(req *http.Request) (*http.Response, error) {
	if req == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, errors.New("request is nil")
	}
	return e.head.Fn(req, e.head.Next)
}
