package messaging

import (
	"fmt"
	"time"
)

func _ExampleIsFinalized() {
	count := 0
	final := IsFinalized(2, time.Second*5, func() bool { return false })
	fmt.Printf("test: IsFinalized() -> [finalized:%v]\n", final)

	final = IsFinalized(2, time.Second*5, func() bool {
		if count == 0 {
			count = 1
			return false
		}
		return true
	})
	fmt.Printf("test: IsFinalized() -> [count:%v] [finalized:%v]\n", count, final)

	//Output:
	//test: IsFinalized() -> [finalized:false]
	//test: IsFinalized() -> [finalized:true]

}

func ExampleIsFinalized_False() {
	final := IsFinalized(2, time.Second*5, func() bool {
		time.Sleep(time.Second * 2)
		return false
	})
	fmt.Printf("test: IsFinalized() -> [finalized:%v]\n", final)

	//Output:
	//test: IsFinalized() -> [finalized:false]

}

func ExampleIsFinalized_True() {
	count := 0
	final := IsFinalized(2, time.Second*5, func() bool {
		time.Sleep(time.Second * 2)
		if count == 1 {
			return true
		}
		count++
		return false
	})
	fmt.Printf("test: IsFinalized() -> [count:%v] [finalized:%v]\n", count, final)

	//Output:
	//test: IsFinalized() -> [count:1] [finalized:true]

}
