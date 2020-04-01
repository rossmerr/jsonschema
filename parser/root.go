package parser

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

type Parse struct {
	Structs map[jsonschema.ID]*document.Document
}

func NewParse() *Parse {
	return &Parse{
		Structs: map[jsonschema.ID]*document.Document{},
	}
}
