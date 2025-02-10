package jsonx

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/aspect"
)

func Marshal(v any) ([]byte, *aspect.Status) {
	if v == nil {
		return nil, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("invalid argument: value is nil"))
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, aspect.NewStatusError(aspect.StatusJsonEncodeError, err)
	}
	return buf, aspect.StatusOK()

}
