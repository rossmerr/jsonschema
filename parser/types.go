package parser

type Types interface {
	Comment() string
	Name() string
	WithFieldTag(string) Types
	WithReference(bool) Types
}
