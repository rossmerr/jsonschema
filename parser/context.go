package parser

import (
	"context"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
)

type SchemaContext struct {
	context.Context
	References   map[string]*jsonschema.Schema
	Package      string
	Tags         tags.FieldTag
	ParentSchema *jsonschema.Schema
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		map[string]*jsonschema.Schema{},
		packageName,
		tags,
		nil,
	}
}

func WrapContext(ctx *SchemaContext, schema *jsonschema.Schema) *SchemaContext {
	return &SchemaContext{
		ctx,
		ctx.References,
		ctx.Package,
		ctx.Tags,
		schema,
	}
}
