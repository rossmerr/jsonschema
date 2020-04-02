package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleInteger(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	return types.NewInteger(name, schema.Description), nil
}
