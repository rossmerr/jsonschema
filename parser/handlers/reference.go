package handlers

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleReference(ctx *parser.SchemaContext, doc parser.Root, name string, schema *jsonschema.Schema) (parser.Component, error) {
	typename, err := ctx.ResolvePointer(schema.Ref, schema.RootSchema())

	if err != nil {
		fmt.Printf("handlereference: reference not found %v\n", schema.Ref)
	}

	return templates.NewReference(name, schema.Description, parser.NewType(typename, parser.Reference)), nil
}
