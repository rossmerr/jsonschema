package document

import (
	"fmt"
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Process func(name string, schema *jsonschema.Schema) (Types, error)
type HandleSchemaFunc func(*DocumentContext, string, *jsonschema.Schema) (Types, error)

type DocumentContext struct {
	References      map[jsonschema.ID]jsonschema.JsonSchema
	implementations map[string][]string
	Package         string
	parentSchema    *jsonschema.RootSchema
	Globals         map[string]Types
	resolve         func(name string, schema jsonschema.JsonSchema) HandleSchemaFunc
}

func NewDocumentContext(packageName string, resolve func(name string, schema jsonschema.JsonSchema) HandleSchemaFunc, references map[jsonschema.ID]jsonschema.JsonSchema, parent *jsonschema.RootSchema) *DocumentContext {
	return &DocumentContext{
		references,
		map[string][]string{},
		packageName,
		parent,
		map[string]Types{},
		resolve,
	}
}

func (ctx *DocumentContext) Parent() *jsonschema.RootSchema {
	return ctx.parentSchema
}

func (ctx *DocumentContext) AddMethods(structname string, methods ...string) {
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

func (ctx *DocumentContext) GetMethods(structname string) []string {
	structname = strings.ToLower(structname)
	return ctx.implementations[structname]
}

func (ctx *DocumentContext) Process(name string, schema jsonschema.JsonSchema) (Types, error) {
	handler := ctx.resolve(name, schema)
	switch v := schema.(type) {
	case *jsonschema.RootSchema:
		return handler(ctx, name, v.Schema)
	case *jsonschema.Schema:
		return handler(ctx, name, v)
	}
	return nil, fmt.Errorf("documentcontext: type of schmema not found %v", schema)
}
