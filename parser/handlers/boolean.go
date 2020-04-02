package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleBoolean(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	return types.NewBoolean(name, schema.Description), nil
}
