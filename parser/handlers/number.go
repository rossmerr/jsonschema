package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleNumber(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	return types.NewNumber(name, schema.Description), nil
}
