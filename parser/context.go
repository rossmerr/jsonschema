package parser

import (
	"context"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
)

type SchemaContext struct {
	context.Context
	References   map[string]*jsonschema.Schema
	Implementations map[string][]string
	Package      string
	Tags         tags.FieldTag
	ParentSchema *jsonschema.Schema
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		map[string]*jsonschema.Schema{},
		map[string][]string{},
		packageName,
		tags,
		nil,
	}
}

func WrapContext(ctx *SchemaContext, schema *jsonschema.Schema) *SchemaContext {
	return &SchemaContext{
		ctx,
		ctx.References,
		ctx.Implementations,
		ctx.Package,
		ctx.Tags,
		schema,
	}
}

func (ctx *SchemaContext) Base() (*jsonschema.Schema, *SchemaContext) {
	if base := ctx.base(); base != nil {
		return base.ParentSchema, base
	}
	return nil, nil
}

func (ctx *SchemaContext) base() *SchemaContext {
	if ctx.ParentSchema == nil {
		return nil
	}

	if ctx.ParentSchema.ID != jsonschema.EmptyString {
		return ctx
	} else {
		if c, ok := ctx.Context.(*SchemaContext); ok {
			return c.base()
		}
	}

	return nil
}
