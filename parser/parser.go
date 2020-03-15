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
	Parse(schemas map[string]*jsonschema.Schema) *Parse
}

type parser struct {
	ctx *SchemaContext
}

func NewParser(ctx context.Context, packageName string) Parser {
	return &parser{
		ctx: NewContext(ctx, packageName, tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})),
	}
}

func (s *parser) Parse(schemas map[string]*jsonschema.Schema) *Parse {
	parse := NewParse()

	for _, schema := range schemas {
		for key, definition := range schema.Definitions {
			s.ctx.Definitions[key] = SchemaToType(s.ctx, key, definition, nil)
		}
	}

	for key, schema := range schemas {
		switch schema.Type() {
		case reflect.Struct:
			parse.Structs[schema.ID] = NewStruct(s.ctx, key, schema, nil)
		case reflect.Interface:
			parse.Interfaces[schema.ID] = NewInterface(schema)
		}
	}

	return parse
}

func SchemaToType(ctx *SchemaContext, key string, schema, parent *jsonschema.Schema) Types {
	switch schema.Type() {
	case reflect.Struct:
		return NewStruct(ctx, key, schema, parent)
	case reflect.Interface:
		return NewInterface(schema)
	case reflect.Array:
		return NewArray(ctx, key, schema, parent)
	case reflect.Int32:
		fallthrough
	case reflect.Float64:
		return NewNumber(ctx, key, schema, parent)
	case reflect.String:
		return NewString(ctx, key, schema, parent)
	case reflect.Ptr:
		return ctx.Definitions[schema.Ref]
	default:
		return nil
	}

	return nil
}
