package parser

import "strconv"

type Kind int

const (
	Boolean Kind = iota
	Enum
	String
	Interger
	Number
	Array
	Reference
	OneOf
	AnyOf
	AllOf
	Object
	RootDocument
	Invalid
)

func (s Kind) String() string {
	for v, k := range kindNames {
		if k == s {
			return v
		}
	}
	return "kind" + strconv.Itoa(int(s))
}

var kindNames = map[string]Kind{
	"boolean":   Boolean,
	"enum":      Enum,
	"string":    String,
	"interger":  Interger,
	"number":    Number,
	"array":     Array,
	"reference": Reference,
	"oneof":     OneOf,
	"anyof":     AnyOf,
	"allof":     AllOf,
	"object":    Object,
	"rootdocument":      RootDocument,
	"invalid":   Invalid,
}
