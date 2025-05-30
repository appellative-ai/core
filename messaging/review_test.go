package messaging

import (
	"fmt"
	"time"
)

func ExampleNewReview() {
	dur := 0
	r := NewReview(dur)
	fmt.Printf("test: NewReview(\"%v\") -> [started:%v] [expired:%v] [dur:%v]\n", dur, r.Started(), r.Expired(), r.duration)

	dur = 1
	r = NewReview(dur)
	fmt.Printf("test: NewReview(\"%v\") -> [started:%v] [expired:%v] [dur:%v]\n", dur, r.Started(), r.Expired(), r.duration)

	dur = 5
	r = NewReview(dur)
	fmt.Printf("test: NewReview(\"%v\") -> [started:%v] [expired:%v] [dur:%v]\n", dur, r.Started(), r.Expired(), r.duration)

	//Output:
	//test: NewReview("0") -> [started:false] [expired:true] [dur:1m0s]
	//test: NewReview("1") -> [started:false] [expired:true] [dur:1m0s]
	//test: NewReview("5") -> [started:false] [expired:true] [dur:5m0s]

}

func ExampleReview_Start() {
	r := newReview(time.Millisecond * 500)
	fmt.Printf("test: NewReview() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	r.Start()
	fmt.Printf("test: Start() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)
	time.Sleep(time.Millisecond * 750)

	fmt.Printf("test: Start() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	//Output:
	//test: NewReview() -> [started:false] [expired:true] [dur:500ms]
	//test: Start() -> [started:true] [expired:false] [dur:500ms]
	//test: Start() -> [started:true] [expired:true] [dur:500ms]

}
