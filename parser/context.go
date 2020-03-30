package parser

import (
	"context"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/tags"
)

type SchemaContext struct {
	context.Context
	References      map[jsonschema.ID]*jsonschema.Schema
	implementations map[string][]string
	Package         string
	Tags            tags.FieldTag
	parentSchema    *jsonschema.Schema
	Globals         map[string]Types
}

func NewContext(ctx context.Context, packageName string, tags tags.FieldTag) *SchemaContext {
	return &SchemaContext{
		ctx,
		map[jsonschema.ID]*jsonschema.Schema{},
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
	if schema.ID.IsNotEmpty() {
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
		switch arr, ok := ctx.implementations[structname]; {
		case !ok:
			arr = []string{}
			fallthrough
		default:
			arr = append(arr, methods...)
			ctx.implementations[structname] = jsonschema.Unique(arr)
		}
	}
}

func (ctx *SchemaContext) GetMethods(structname string) []string {
	structname = strings.ToLower(structname)
	return ctx.implementations[structname]
}
