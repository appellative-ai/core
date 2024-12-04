package jsonx

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/core"
)

func Marshal(v any) ([]byte, *core.Status) {
	if v == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: value is nil"))
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, core.NewStatusError(core.StatusJsonEncodeError, err)
	}
	return buf, core.StatusOK()

}
