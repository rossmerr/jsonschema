package parser

import "github.com/RossMerr/jsonschema"

type Component interface {
	Name() string
}

type Root interface {
	WithPackageName(packagename string)
	Globals()    map[string]Component
	AddImport(value string)
	Root() *jsonschema.Schema
}

type Field interface {
	WithFieldTag(string) Field
	FieldTag() string
	WithReference(bool) Field
}

type Receiver interface {
	WithMethods(methods ...*Method)
}