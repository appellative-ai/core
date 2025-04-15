package accesslogtest

import (
	"fmt"
	"github.com/behavioral-ai/core/accessdata"
)

// DebugOutputHandler - output to stdio
type DebugOutputHandler struct{}

func (DebugOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	fmt.Printf("%v\n", formatter.Format(items, data))
}

// TestOutputHandler - special use case of DebugOutputHandler for testing examples
type TestOutputHandler struct{}

func (TestOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	fmt.Printf("test: Write() -> [%v]\n", formatter.Format(items, data))
}
