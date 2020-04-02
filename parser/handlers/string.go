package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleString(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	return templates.NewString(name, schema.Description), nil
}
