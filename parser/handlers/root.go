package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleRoot(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	t, err := doc.Process(name, schema)
	return types.NewRoot(schema.Description, t), err
}
