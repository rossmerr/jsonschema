package parser

import (
	"context"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
)

type SchemaContext struct {
	context.Context
	Interfaces map[string]*jsonschema.Schema
	References map[string]*jsonschema.Schema
	Package    string
	Tags       tags.FieldTag
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		map[string]*jsonschema.Schema{},
		map[string]*jsonschema.Schema{},
		packageName,
		tags,
	}
}
