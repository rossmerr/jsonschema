package parser

import (
	"context"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
)

type SchemaContext struct {
	context.Context
	References      map[string]*jsonschema.Schema
	implementations map[string][]string
	Package         string
	Tags            tags.FieldTag
	parentSchema    *jsonschema.Schema
	Globals         map[string]Types
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		map[string]*jsonschema.Schema{},
		map[string][]string{},
		packageName,
		tags,
		nil,
		map[string]Types{},
	}
}

func (ctx *SchemaContext) SetParent(schema *jsonschema.Schema) *SchemaContext {
	if schema == nil {
		return ctx
	}
	if schema.ID != jsonschema.EmptyString {
		ctx.parentSchema = schema
	}
	return ctx
}

func (ctx *SchemaContext) Parent() *jsonschema.Schema {
	return ctx.parentSchema
}

func (ctx *SchemaContext) AddMethods(structname string, methods ...string) {
	if structname != jsonschema.EmptyString {
		structname = strings.ToLower(structname)
		arr := ctx.implementations[structname]
		if arr == nil {
			arr = []string{}
		}
		arr = append(arr, methods...)
		ctx.implementations[structname] = arr
	}
}

func (ctx *SchemaContext) GetMethods(structname string) []string {
	structname = strings.ToLower(structname)
	return ctx.implementations[structname]
}
