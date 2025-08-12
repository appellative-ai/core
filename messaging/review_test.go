package messaging

import (
	"fmt"
	"time"
)

func ExampleNewReview() {
	r := NewReview()
	fmt.Printf("test: NewReview() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	r = NewReview()
	fmt.Printf("test: NewReview() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	r = NewReview()
	fmt.Printf("test: NewReview() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	//Output:
	//test: NewReview() -> [started:false] [expired:false] [dur:0s]
	//test: NewReview() -> [started:false] [expired:false] [dur:0s]
	//test: NewReview() -> [started:false] [expired:false] [dur:0s]

}

func ExampleReview_Start() {
	r := NewReview()
	fmt.Printf("test: NewReview() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	r.Start(time.Millisecond * 500)
	fmt.Printf("test: Start() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)
	time.Sleep(time.Millisecond * 750)

	fmt.Printf("test: Start() -> [started:%v] [expired:%v] [dur:%v]\n", r.Started(), r.Expired(), r.duration)

	//Output:
	//test: NewReview() -> [started:false] [expired:false] [dur:0s]
	//test: Start() -> [started:true] [expired:false] [dur:500ms]
	//test: Start() -> [started:false] [expired:true] [dur:500ms]

}
