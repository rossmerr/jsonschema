package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleArray(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	arrType := string(schema.Items.Type)
	if schema.Items.Ref.IsNotEmpty() {
		arrType = schema.Items.Ref.ToTypename()
	}

	return types.NewArray(name, schema.Description, arrType), nil
}
