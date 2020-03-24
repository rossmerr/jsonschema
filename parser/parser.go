package parser

import (
	"context"

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
		switch schema.Type {
		case jsonschema.Object:
			filename := schema.ID.Filename()
			anonymousStruct := NewAnonymousStruct(s.ctx, filename, schema, nil)

			definitions := make([]*Definition, 0)
			for typename, def := range schema.Definitions {
				definition := NewDefinition(WrapContext(s.ctx, schema), typename, def)
				definitions = append(definitions, definition)
			}
			for typename, def := range schema.Defs {
				definition := NewDefinition(WrapContext(s.ctx, schema), typename, def)
				definitions = append(definitions, definition)
			}
			parse.Structs[schema.ID] = NewStruct(s.ctx, anonymousStruct,definitions, filename)
		}
	}

	return parse
}

func (s *parser) buildReferences(schemas map[jsonschema.ID]*jsonschema.Schema) {
	for _, schema := range schemas {
		key := schema.ID.Base()
		s.ctx.References[key] = schema
	}
}

func SchemaToType(ctx *SchemaContext, field string, schema *jsonschema.Schema, required []string) Types {
	field = jsonschema.Fieldname(field)
	switch schema.Type {
	case jsonschema.Boolean:
		return NewBoolean(ctx, field, schema, required)
	case jsonschema.String:
		// if schema.IsEnum() {
		// 	return NewEnum(ctx, field, schema, required)
		// }
		return NewString(ctx, field, schema, required)
	case jsonschema.Integer:
		return NewInteger(ctx, field, schema, required)
	case jsonschema.Number:
		return NewNumber(ctx, field, schema, required)
	case jsonschema.Array:
		return NewArray(ctx, field, schema, required)
	case jsonschema.Object:
		fallthrough
	default:
		if RequiesInterface(schema) {
			t := NewInterfaceReference(ctx,  field, schema)
			return t
		}

		if schema.Ref != jsonschema.EmptyString {
			_, typename, _ := ResolvePointer(ctx, schema.Ref)
			t := NewReference(typename, field)
			return t
		}

		return NewAnonymousStruct(ctx, field, schema, required)
	}
}

func RequiesInterface(s *jsonschema.Schema) bool {
	if s.OneOf != nil && len(s.OneOf) > 0 {
		return true
	}
	if s.AnyOf != nil && len(s.AnyOf) > 0 {
		return true
	}
	if s.AllOf != nil && len(s.AllOf) > 0 {
		return true
	}

	return false
}

