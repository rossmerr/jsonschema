package handlers

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleReference(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	typename, err := doc.ResolvePointer(schema.Ref)

	if err != nil {
		fmt.Printf("handlereference: reference not found %v\n", schema.Ref)
	}

	return templates.NewReference(name, schema.Description, typename), nil
}
