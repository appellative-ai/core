package accessdata

var (
	operators2 = []Operator{
		{Name: "start-time", Value: StartTimeOperator},
		{Name: "duration-ms", Value: DurationOperator},
		{Name: "traffic", Value: TrafficOperator},
		//{Name: "route-name", Value: RouteNameOperator},

		{Name: "region", Value: OriginRegionOperator},
		{Name: "zone", Value: OriginZoneOperator},
		{Name: "sub-zone", Value: OriginSubZoneOperator},
		{Name: "service", Value: OriginServiceOperator},
		{Name: "instance-id", Value: OriginInstanceIdOperator},

		{Name: "method", Value: RequestMethodOperator},
		{Name: "url", Value: RequestUrlOperator},
		{Name: "host", Value: RequestHostOperator},
		{Name: "path", Value: RequestPathOperator},
		{Name: "protocol", Value: RequestProtocolOperator},
		{Name: "request-id", Value: RequestIdOperator},
		{Name: "forwarded", Value: RequestForwardedForOperator},

		{Name: "status-code", Value: ResponseStatusCodeOperator},

		{Name: "timeout-ms", Value: TimeoutDurationOperator},
		{Name: "rate-limit", Value: RateLimitOperator},
		{Name: "rate-burst", Value: RateBurstOperator},
		//{Name: "retry", Value: RetryOperator},
		{Name: "proxy", Value: ProxyOperator},
		{Name: "status-flags", Value: StatusFlagsOperator},
	}
)
