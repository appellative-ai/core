package accesslog

import (
	"github.com/behavioral-ai/core/accessdata"
	"log"
)

// OutputHandler - template parameter for accesslog output
type OutputHandler interface {
	Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter)
}

// NilOutputHandler - no output
type NilOutputHandler struct{}

func (NilOutputHandler) Write(_ []accessdata.Operator, _ *accessdata.Entry, _ accessdata.Formatter) {
}

// LogOutputHandler - output to accesslog
type LogOutputHandler struct{}

func (LogOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	log.Println(formatter.Format(items, data))
}
