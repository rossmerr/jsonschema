package jsonschema

import "reflect"

type Kind string

const (
	String  Kind = "string"
	Object  Kind = "object"
	Array   Kind = "array"
	Integer Kind = "integer"
	Number  Kind = "number"
	Boolean Kind = "boolean"
	Null    Kind = "null"
)

func (s Kind) String() string {
	return string(s)
}

func (s Kind) ToKind() reflect.Kind {
	switch s {
	case String:
		return reflect.String
	case Object:
		return reflect.Struct
	case Array:
		return reflect.Slice
	case Integer:
		return reflect.Int32
	case Number:
		return reflect.Float64
	case Boolean:
		return reflect.Bool
	case Null:
		return reflect.Invalid
	default:
		return reflect.Struct
	}
}
