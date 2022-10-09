package common

import (
	"reflect"
	"strconv"
)

func Any(value any) string {
	if valueStruct, ok := value.(reflect.Value); ok {
		return atomConvert(valueStruct)
	}
	return atomConvert(reflect.ValueOf(value))
}

func atomConvert(value reflect.Value) string {
	kind := value.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.String:
		// return strconv.Quote(value.String())
		return value.String()
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.Float32:
		return strconv.FormatFloat(value.Float(), 'e', 1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'e', 1, 64)
	case reflect.Slice, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr: // reference type
		return "reference type" + kind.String() + " 0x" + strconv.FormatUint(uint64(value.Pointer()), 16)
	default: // for Invalid, Array, Struct, Interface
		return "Unsupport type" + kind.String()
	}
}
