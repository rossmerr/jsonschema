package parser

import "github.com/RossMerr/jsonschema"

type Parse struct {
	Structs    map[jsonschema.ID]*Struct
	Interfaces map[jsonschema.ID]*Interface
}

func NewParse() *Parse {
	return &Parse{
		Structs:    map[jsonschema.ID]*Struct{},
		Interfaces: map[jsonschema.ID]*Interface{},
	}
}
