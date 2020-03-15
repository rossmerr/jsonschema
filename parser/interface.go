package parser

import "github.com/RossMerr/jsonschema"

type Interface struct {
	Comment string
	Name    string
}

func NewInterface(schema *jsonschema.Schema) *Interface {
	return &Interface{
		Comment: schema.Description,
		Name:    string(schema.ID),
	}
}

func (s *Interface) types() {}
