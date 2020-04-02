package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleBoolean(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	return types.NewBoolean(name, schema.Description), nil
}
