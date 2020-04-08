package parser

type Types interface {
	Comment() string
	Name() string
	WithFieldTag(string) Types
	FieldTag() string
	WithReference(bool) Types
	WithMethods(methods ...*Method) Types
}

type Component interface {
	Name() string
}

type Root interface {
	WithPackageName(packagename string)
}
