package jsonx

import (
	"encoding/json"
	"errors"
)

func Marshal(v any) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid argument: value is nil") //aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("invalid argument: value is nil"))
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err //aspect.NewStatusError(aspect.StatusJsonEncodeError, err)
	}
	return buf, nil //aspect.StatusOK()

}
