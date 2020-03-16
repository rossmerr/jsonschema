package parser

import (
	"context"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
)

type SchemaContext struct {
	context.Context
	Refer   map[jsonschema.ID]*jsonschema.Schema
	Package string
	Tags    tags.FieldTag
	ImplementInterface map[jsonschema.ID][]string
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		map[jsonschema.ID]*jsonschema.Schema{},
		packageName,
		tags,
		map[jsonschema.ID][]string{},
	}
}
