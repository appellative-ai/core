package messaging

import "fmt"

func ExampleNewReview() {
	dur := ""
	r := NewReview(dur)
	fmt.Printf("test: NewReview(\"%v\") -> [started:%v] [scheduled:%v] [expired:%v] [dur:%v]\n", dur, r.Started(), r.Scheduled(), r.Expired(), r.duration)

	dur = "500ms"
	r = NewReview(dur)
	fmt.Printf("test: NewReview(\"%v\") -> [started:%v] [scheduled:%v] [expired:%v] [dur:%v]\n", dur, r.Started(), r.Scheduled(), r.Expired(), r.duration)

	dur = "5m"
	r = NewReview(dur)
	fmt.Printf("test: NewReview(\"%v\") -> [started:%v] [scheduled:%v] [expired:%v] [dur:%v]\n", dur, r.Started(), r.Scheduled(), r.Expired(), r.duration)

	//Output:
	//test: NewReview("") -> [started:false] [scheduled:false] [expired:true] [dur:0s]
	//test: NewReview("500ms") -> [started:false] [scheduled:true] [expired:true] [dur:1m0s]
	//test: NewReview("5m") -> [started:false] [scheduled:true] [expired:true] [dur:5m0s]

}
