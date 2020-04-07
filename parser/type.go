package parser

type Type struct {
	Kind Kind
	Name string
}

func NewType(name string, kind Kind) *Type {
	return &Type{
		Name: name,
		Kind: kind,
	}
}
