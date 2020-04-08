package jsonschema

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

