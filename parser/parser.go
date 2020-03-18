package parser

import (
	"context"
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
			parse.Structs[schema.ID] = NewStruct(s.ctx, NewAnonymousStruct(s.ctx, schema, nil))
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


func SchemaToType(ctx *SchemaContext, key jsonschema.ID, schema, parent *jsonschema.Schema) Types {
	switch schema.Type() {
	case reflect.Struct:
		schema.ID = key
		return NewAnonymousStruct(ctx, schema, parent)
	case reflect.Interface:
		schema.ID = key
		return NewInterface(ctx, schema, parent)
	case reflect.Array:
		return NewArray(ctx, key, schema, parent)
	case reflect.Int32:
		fallthrough
	case reflect.Float64:
		return NewNumber(ctx, key, schema, parent)
	case reflect.String:
		return NewString(ctx, key, schema, parent)
	case reflect.Ptr:
		ref := ctx.References[schema.Ref.Base()]
		return SchemaToType(ctx, key, ref, parent)
	default:
		return NewAnonymousStruct(ctx, schema, parent)
	}

	return nil
}
