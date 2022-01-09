package internal

import (
	"github.com/decimalbell/akv"
)

type valueType byte

const (
	valueTypeBytes valueType = 0
	valueTypeList  valueType = 1
	valueTypeSet   valueType = 2
	valueTypeZSet  valueType = 3
	valueTypeHash  valueType = 4
)

type value struct {
	typ valueType
	val interface{}
}

func newBytesValue(val []byte) *value {
	return &value{
		typ: valueTypeBytes,
		val: val,
	}
}

func (v *value) bytes() ([]byte, error) {
	if v.typ != valueTypeBytes {
		return nil, akv.ErrWrongType
	}

	val, ok := v.val.([]byte)
	if !ok {
		return nil, akv.ErrWrongType
	}
	return val, nil
}
