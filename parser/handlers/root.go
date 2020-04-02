package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleRoot(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	// Stops the HandleRoot getting in a loop
	schema.Parent = &jsonschema.Schema{}

	t, err := doc.Process(name, schema)
	if err != nil {
		return nil, err
	}
	return templates.NewRoot(schema.Description, t), nil
}
