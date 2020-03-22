package parser

import "github.com/RossMerr/jsonschema"

type Definition struct {
	*AnonymousStruct
}

func NewDefinition(ctx *SchemaContext, typename string, schema *jsonschema.Schema) *Definition {
	anonymousStruct := NewAnonymousStruct(ctx, typename, schema, nil)

	return &Definition{
		anonymousStruct,
	}
}