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

type Value struct {
	typ valueType
	val interface{}
}

func newBytesValue(val []byte) *Value {
	return &Value{
		typ: valueTypeBytes,
		val: val,
	}
}

func newHashValue() *Value {
	return &Value{
		typ: valueTypeHash,
		val: newMap(),
	}
}

func newSetValue() *Value {
	return &Value{
		typ: valueTypeSet,
		val: newSet(),
	}
}

func (v *Value) bytes() ([]byte, error) {
	if v.typ != valueTypeBytes {
		return nil, akv.ErrWrongType
	}

	val, ok := v.val.([]byte)
	if !ok {
		return nil, akv.ErrWrongType
	}
	return val, nil
}

func (v *Value) hash() (Map, error) {
	if v.typ != valueTypeHash {
		return nil, akv.ErrWrongType
	}

	val, ok := v.val.(Map)
	if !ok {
		return nil, akv.ErrWrongType
	}
	return val, nil
}

func (v *Value) set() (Set, error) {
	if v.typ != valueTypeSet {
		return nil, akv.ErrWrongType
	}

	val, ok := v.val.(Set)
	if !ok {
		return nil, akv.ErrWrongType
	}
	return val, nil
}
