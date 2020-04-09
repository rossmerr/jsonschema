package parser

type Component interface {
	Name() string
}

type Root interface {
	Component
	Globals() map[string]Component
	AddImport(value string)
}

type Field interface {
	Component
	WithFieldTag(string) Field
	FieldTag() string
	WithReference(bool) Field
}

type Receiver interface {
	Field
	WithMethods(methods ...*Method)
}
