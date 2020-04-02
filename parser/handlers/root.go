package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleRoot(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	schema.Parent = &jsonschema.Schema{}
	t, err := doc.Process(name, schema)
	return types.NewRoot(schema.Description, t), err
}
