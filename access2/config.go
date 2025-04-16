package access2

import (
	"encoding/json"
	"errors"
)

// InitEgressOperators - allows configuration of access attributes for egress traffic
/*
func InitEgressOperators(config []accessdata.Operator) error {
	var err error
	egressOperators, err = accessdata.InitOperators(config)
	return err
}

*/

// LoadOperators - load operators from file
func LoadOperators(read func() ([]byte, error)) error {
	if read == nil {
		return errors.New("invalid argument: ReadConfig function is nil")
	}
	buf, err0 := read()
	if err0 != nil {
		return err0
	}
	var ops []Operator

	err := json.Unmarshal(buf, &ops)
	if err != nil {
		return err
	}
	ops, err = InitOperators(ops)
	if err == nil {
		defaultOperators = ops
	}
	return err
}
