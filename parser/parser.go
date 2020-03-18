package parser

import (
	"context"
	"fmt"
	"reflect"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
	"github.com/RossMerr/jsonschema/tags/json"
	"github.com/RossMerr/jsonschema/tags/validate"
)

type Parser interface {
	Parse(schemas map[jsonschema.ID]*jsonschema.Schema) *Parse
}

type parser struct {
	ctx *SchemaContext
}

func NewParser(ctx context.Context, packageName string) Parser {
	return &parser{
		ctx: NewContext(ctx, packageName, tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})),
	}
}

func (s *parser) Parse(schemas map[jsonschema.ID]*jsonschema.Schema) *Parse {
	parse := NewParse()

	s.buildReferences(schemas)

	for _, schema := range schemas {
		switch schema.Type() {
		case reflect.Struct:
			filename := schema.ID.Filename()
			parse.Structs[schema.ID] = NewStruct(s.ctx, NewAnonymousStruct(s.ctx, filename, schema, nil), filename)
		}
	}

	return parse
}

func (s *parser)buildReferences(schemas map[jsonschema.ID]*jsonschema.Schema) {
	for _, schema := range schemas {
		key := schema.ID.Base()
		s.ctx.References[key] = schema
	}
}


func SchemaToType(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string ) Types {
	typename = jsonschema.Typename(typename)
	switch schema.Type() {
	case reflect.Struct:
		return NewAnonymousStruct(ctx, typename, schema, required)
	case reflect.Interface:
		return NewInterface(ctx, typename, schema, required)
	case reflect.Array:
		return NewArray(ctx, typename, schema, required)
	case reflect.Int32:
		fallthrough
	case reflect.Float64:
		return NewNumber(ctx, typename, schema, required)
	case reflect.String:
		return NewString(ctx, typename, schema, required)
	case reflect.Bool:
		return NewBoolean(ctx, typename, schema, required)
	case reflect.Ptr:
		if ref, ok := ctx.References[schema.Ref.Base()]; ok {
			return SchemaToType(ctx, typename, ref, required)
		} else {
			panic(fmt.Errorf("Reference not found! '%v'", schema.Ref.Base()))
		}
	default:
		return NewAnonymousStruct(ctx, typename, schema, required)
	}
}
