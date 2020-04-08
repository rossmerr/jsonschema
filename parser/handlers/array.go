package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleArray(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	arrType := string(schema.Items.Type)
	if schema.Items.Ref.IsNotEmpty() {
		arrType = schema.Items.Ref.ToKey()
	}

	return templates.NewArray(name, schema.Description, arrType), nil
}
