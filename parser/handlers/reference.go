package handlers

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleReference(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	typename, err := types.ResolvePointer(doc, schema.Ref)

	if err != nil {
		fmt.Printf("handlereference: reference not found %v\n", schema.Ref)
	}

	return types.NewReference(name, schema.Description, typename), nil
}
