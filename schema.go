package jsonschema

import (
	"github.com/RossMerr/jsonschema/types"
)

type Schema struct {
	ID string
	Document    types.Document
	Definitions *SchemaReferences
	Config      Config
}
