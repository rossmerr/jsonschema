package parser

import "github.com/RossMerr/jsonschema"

type Definition struct {
	*AnonymousStruct
}

func NewDefinition(ctx *SchemaContext, typename string, schema *jsonschema.Schema) *Definition {
	typename = jsonschema.Structname(typename)
	anonymousStruct := NewAnonymousStruct(ctx, typename, schema, nil)
	arr := ctx.Implementations[typename]
	if len(arr) > 0 {
		anonymousStruct.Method = arr[0]
	}
	return &Definition{
		anonymousStruct,
	}
}