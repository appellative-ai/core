package access

/*
func ExampleDefault_Host() {
	start := time.Now().UTC()
	SetOrigin(Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	req, _ := http.NewRequest("GET", "https://www.google.com/search?q=test", nil)
	req.Header.Add(XRequestId, "123-456")
	//fmt.Printf("test: NewRequest() -> [err:%v] [req:%v]\n", err, req != nil)
	resp := http.Response{StatusCode: http.StatusOK}
	resp.Header = make(http.Header)
	resp.Header.Add(ContentEncoding, "gzip")
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), "", req, &resp, Controller{Timeout: -1})

	fmt.Printf("test: Default-Host() -> %v\n", "success")

	//Output:
	//test: Default-Host() -> success

}

func ExampleDefault_Domain() {
	start := time.Now().UTC()
	values := make(url.Values)
	values.Add("region", "*")
	values.Add("zone", "texas")

	req, _ := http.NewRequest("select", "https://github.com/advanced-go/example-domain/activity:v1/entry?"+BuildQuery(values), nil)
	req.Header.Add(XRequestId, "123-456")
	//fmt.Printf("test: NewRequest() -> [err:%v] [req:%v]\n", err, req != nil)
	resp := http.Response{StatusCode: http.StatusOK}
	logTest(IngressTraffic, start, time.Since(start), "", req, &resp, Controller{Timeout: -1})

	fmt.Printf("test: Default-Domain() -> %v\n", "success")

	//Output:
	//test: Default-Domain() -> success

}

func ExampleDefault_Access_Request_Status() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	req := RequestImpl{Method: http.MethodPut, Url: "https://www.google.com/search?q=test", Header: h}
	start := time.Now().UTC()
	SetOrigin(Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusNotFound
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), "", req, resp, Controller{Timeout: -1})

	fmt.Printf("test: Default-Access-Request-Status() -> %v\n", "success")

	//Output:
	//test: Default-Access-Request-Status() -> success

}

func ExampleDefault_Access_Request_Status_Code() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	req := RequestImpl{Method: http.MethodPut, Url: "https://www.google.com/search?q=test", Header: h}
	start := time.Now().UTC()
	SetOrigin(Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), "", req, resp, Controller{Timeout: -1})

	fmt.Printf("test: Default-Access-Request-Status-Code() -> %v\n", "success")

	//Output:
	//test: Default-Access-Request-Status-Code() -> success

}

func ExampleDefault_Threshold_Duration() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	req := RequestImpl{Method: http.MethodPut, Url: "https://www.google.com/search?q=test", Header: h}
	start := time.Now().UTC()
	SetOrigin(Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), "", req, resp, Controller{Timeout: time.Second * 4})

	fmt.Printf("test: Default-Threshold-Duration() -> %v\n", "success")

	//Output:
	//test: Default-Threshold-Duration() -> success

}

func ExampleDefault_Threshold_Int() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	req := RequestImpl{Method: http.MethodPut, Url: "https://www.google.com/search?q=test", Header: h}
	start := time.Now().UTC()
	SetOrigin(Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), "", req, resp, Controller{Timeout: -1, RateLimit: "345"})

	fmt.Printf("test: Default-Threshold-Int() -> %v\n", "success")

	//Output:
	//test: Default-Threshold-Int() -> success

}

func ExampleDefault_Threshold_Deadline() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	req := RequestImpl{Method: http.MethodPut, Url: "https://www.google.com/search?q=test", Header: h}
	start := time.Now().UTC()
	SetOrigin(Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	//ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), "", req, resp, Controller{})

	fmt.Printf("test: Default-Threshold-Int() -> %v\n", "success")

	//Output:
	//test: Default-Threshold-Int() -> success

}

func logTest(traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds []any) {
	Log(traffic, start, duration, route, req, resp, thresholds)
}


*/
