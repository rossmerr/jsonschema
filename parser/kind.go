package parser

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
	RootObject
)
