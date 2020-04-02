package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleRoot(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	schema.Parent = &jsonschema.Schema{}
	t, err := doc.Process(name, schema)
	return types.NewRoot(schema.Description, t), err
}
