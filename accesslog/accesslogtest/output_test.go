package accesslogtest

import (
	"fmt"
	"github.com/behavioral-ai/core/accessdata"
	"github.com/behavioral-ai/core/accesslog"
)

func ExampleOutputHandler() {
	fmt.Printf("test: Output[DebugOutputHandler,data.JsonFormatter](operators,data)\n")
	ops := []accessdata.Operator{{"error", "message"}}
	logTest[DebugOutputHandler, accessdata.JsonFormatter](ops, accessdata.NewEmptyEntry())

	fmt.Printf("test: Output[TestOutputHandler,data.JsonFormatter](nil,nil)\n")
	logTest[TestOutputHandler, accessdata.JsonFormatter](nil, nil)

	fmt.Printf("test: Output[TestOutputHandler,data.JsonFormatter](ops,data)\n")
	logTest[TestOutputHandler, accessdata.JsonFormatter](ops, accessdata.NewEmptyEntry())

	//Output:
	//test: Output[DebugOutputHandler,data.JsonFormatter](operators,data)
	//{"error":"message"}
	//test: Output[TestOutputHandler,data.JsonFormatter](nil,nil)
	//test: Write() -> [{}]
	//test: Output[TestOutputHandler,data.JsonFormatter](ops,data)
	//test: Write() -> [{"error":"message"}]

}

func logTest[O accesslog.OutputHandler, F accessdata.Formatter](items []accessdata.Operator, data *accessdata.Entry) {
	var o O
	var f F
	o.Write(items, data, f)
}
