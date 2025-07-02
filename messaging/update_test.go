package messaging

import "fmt"

func ExampleUpdateMap() {
	UpdateMap("", nil, nil)
	fmt.Printf("test: UpdateMap() -> nil update fn\n")

	fn := func(cfg map[string]string) {
		fmt.Printf("test: UpdateMap() -> %v\n", cfg)
	}
	UpdateMap("", fn, nil)
	fmt.Printf("test: UpdateMap() -> nil message\n")

	m := NewMessage(ChannelControl, "test-message")
	UpdateMap("", fn, m)
	fmt.Printf("test: UpdateMap() -> invalid content type\n")

	cfg := map[string]string{"name": "value"}
	m = NewMapMessage(cfg)
	UpdateMap("", fn, m)

	//Output:
	//test: UpdateMap() -> nil update fn
	//test: UpdateMap() -> nil message
	//test: UpdateMap() -> invalid content type
	//test: UpdateMap() -> map[name:value]

}

func ExampleUpdateReview() {
	UpdateReview("", nil, nil)
	fmt.Printf("test: UpdateReview() -> nil Review\n")

	review := NewReview(5)
	UpdateReview("", &review, nil)
	fmt.Printf("test: UpdateReview() -> nil message\n")

	m := NewMessage(ChannelControl, "test-message")
	UpdateReview("", &review, m)
	fmt.Printf("test: UpdateReview() -> invalid content type\n")

	before := review
	m = NewReviewMessage(NewReview(10))
	UpdateReview("", &review, m)
	fmt.Printf("test: UpdateReview() -> [original:%v] [updated:%v]\n", before.duration, review.duration)

	//Output:
	//test: UpdateReview() -> nil Review
	//test: UpdateReview() -> nil message
	//test: UpdateReview() -> invalid content type
	//test: UpdateReview() -> [original:5m0s] [updated:10m0s]

}
