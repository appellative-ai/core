package httpx

type ChainableExchange func(next Exchange) Exchange

func NewPipeline(ex ...ChainableExchange) Exchange {
	if len(ex) == 0 {
		return nil
	}
	var head Exchange

	for i := len(ex) - 1; i >= 0; i-- {
		f := ex[i]
		if i == len(ex)-1 {
			head = f(nil)
		} else {
			head = f(head)
		}
	}
	return head
}
