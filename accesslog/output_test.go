package accesslog

import (
	"fmt"
	"github.com/behavioral-ai/core/accessdata"
)

func ExampleOutputHandler() {
	fmt.Printf("test: Output[NilOutputHandler,data.TextFormatter](nil,nil)\n")
	logTest[NilOutputHandler, accessdata.TextFormatter](nil, nil)

	ops := []accessdata.Operator{{"error", "message"}}

	fmt.Printf("test: Output[LogOutputHandler,data.JsonFormatter](ops,data)\n")
	logTest[LogOutputHandler, accessdata.JsonFormatter](ops, accessdata.NewEmptyEntry())

	//Output:
	//{"error":"message"}
	//test: Output[NilOutputHandler,data.TextFormatter](nil,nil)
	//test: Output[DebugOutputHandler,data.JsonFormatter](operators,data)

}

func logTest[O OutputHandler, F accessdata.Formatter](items []accessdata.Operator, data *accessdata.Entry) {
	var o O
	var f F
	o.Write(items, data, f)
}
