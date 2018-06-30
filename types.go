package steppingdb

import (
	"fmt"
	"reflect"
)

type DeleteType int

var Delete *DeleteType

type ResizeType int

var Resize *DeleteType

type mapDiff struct {
	k      string
	v      Value
	new    bool
	delete bool
}

type sliceDiff struct {
	at int // -1 append
	v  Value
}

type TypeClass int

const (
	Base TypeClass = iota
	Map
	Array
)

func CheckType(value Value, shouldbe TypeClass) error {
	switch shouldbe {
	case Base:
		return checkBase(value)
	case Map:
		return checkMap(value)
	case Array:
		return checkSlice(value)
	}
}

func checkBase(value Value) error {
	switch value.(type) {
	case float64:
	case int64:
	case uint64:
	case int:
	case []byte:
	case string:
	default:
		return fmt.Errorf("type not supported: %v", reflect.TypeOf(value).String())
	}
	return nil
}

func checkMap(value Value) error {
	if reflect.TypeOf(value).Kind() == reflect.Map {
		return nil
	} else {
		return fmt.Errorf("type not a map: %v", reflect.TypeOf(value).String())
	}
}

func checkSlice(value Value) error {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		return nil
	} else {
		return fmt.Errorf("type not a slice: %v", reflect.TypeOf(value).String())
	}
}
