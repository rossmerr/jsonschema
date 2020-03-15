package parser

import (
	"context"

	"github.com/RossMerr/jsonschema/tags"
)

type SchemaContext struct {
	context.Context
	Definitions Definitions
	Package     string
	Tags        tags.FieldTag
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		Definitions{},
		packageName,
		tags,
	}
}
