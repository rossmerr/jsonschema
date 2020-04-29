package js

type DataType int

const (
	Number  DataType = 0
	Integer DataType = 1
	String  DataType = 2
	Boolean DataType = 3
	Array   DataType = 4
	Object  DataType = 5
	Null    DataType = 6
)

var DataType_name = map[int32]string{
	0: "number",
	1: "integer",
	2: "string",
	3: "boolean",
	4: "array",
	5: "object",
	6: "null",
}

var DataType_value = map[string]int32{
	"number":  0,
	"integer": 1,
	"string":  2,
	"boolean": 3,
	"array":   4,
	"object":  5,
	"null":    6,
}

func (x DataType) String() string {
	return DataType_name[int32(x)]
}
