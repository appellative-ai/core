package messagingtest

import (
	"fmt"
	"time"
)

func ExampleNewTestSpanner() {
	s := NewTestSpanner(time.Second*2, time.Second*5)

	fmt.Printf("test: NewTestSpanner() [dur:%v]\n", s.Span())
	fmt.Printf("test: NewTestSpanner() [dur:%v]\n", s.Span())
	fmt.Printf("test: NewTestSpanner() [dur:%v]\n", s.Span())
	fmt.Printf("test: NewTestSpanner() [dur:%v]\n", s.Span())

	//Output:
	//test: NewTestSpanner() [dur:2s]
	//test: NewTestSpanner() [dur:5s]
	//test: NewTestSpanner() [dur:2s]
	//test: NewTestSpanner() [dur:5s]

}
