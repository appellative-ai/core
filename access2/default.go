package access2

var (
	defaultOperators []Operator
)

func init() {
	defaultOperators, _ = CreateOperators([]string{TrafficOperator,
		StartTimeOperator,
		DurationOperator,
		RouteOperator,
		RequestMethodOperator,
		RequestUrlOperator,
		ResponseStatusCodeOperator,
		ResponseCachedOperator,
		ResponseContentEncodingOperator,
		ResponseBytesReceivedOperator,
		TimeoutDurationOperator,
		RateLimitOperator,
		RedirectOperator,
	})
}
