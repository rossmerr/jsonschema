package parser

type Parameter struct {
	Name string
	Type string
}

func NewParameter(name, typename string) *Parameter {
	return &Parameter{
		Name: name,
		Type: typename,
	}
}
