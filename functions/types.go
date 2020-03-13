package functions

import "reflect"

func IsStruct(value reflect.Kind) bool {
	return value == reflect.Struct
}

func IsInterface(value reflect.Kind) bool {
	return value == reflect.Interface
}

func IsArray(value reflect.Kind) bool {
	return value == reflect.Array
}

func IsString(value reflect.Kind) bool {
	return value == reflect.String
}

func IsNumber(value reflect.Kind) bool {
	return value == reflect.Int32 || value == reflect.Float64
}

func IsPointer(value reflect.Kind) bool {
	return value == reflect.Ptr
}
