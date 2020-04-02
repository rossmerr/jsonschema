package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleBoolean(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	return templates.NewBoolean(name, schema.Description), nil
}
