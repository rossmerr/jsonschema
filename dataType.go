package jsonschema

// DataType specifies the data type for a schema
type DataType string

// JSON Schema defines the following basic types
const (
	// String type is used for strings of text. It may contain Unicode characters
	String  DataType = "string"
	// Object are the mapping type in JSON. They map “keys” to “values”. In JSON, the “keys” must always be strings. Each of these pairs is conventionally referred to as a “property”
	Object  DataType = "object"
	// Array are used for ordered elements. In JSON, each element in an array may be of a different type
	Array   DataType = "array"
	// Integer type is used for integral numbers
	Integer DataType = "integer"
	// Number type is used for any numeric type, either integers or floating point numbers
	Number  DataType = "number"
	// Boolean type matches only two special values: true and false. Note that values that evaluate to true or false, such as 1 and 0, are not accepted by the schema
	Boolean DataType = "boolean"
	// Null type is generally used to represent a missing value. When a schema specifies a type of null, it has only one acceptable value: null
	Null    DataType = "null"
)

// String returns the string representation of the DataType
func (s DataType) String() string {
	return string(s)
}
