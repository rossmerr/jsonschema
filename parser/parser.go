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
			for typename, def := range schema.Definitions  {
				definition := NewDefinition(WrapContext(s.ctx, schema), typename, def)
				definitions = append(definitions, definition)
			}
			for typename, def := range schema.Defs  {
				definition := NewDefinition(WrapContext(s.ctx, schema), typename, def)
				definitions = append(definitions, definition)
			}
			parse.Structs[schema.ID] = NewStruct(s.ctx, anonymousStruct, definitions, filename)
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
	switch schema.Type {
	case jsonschema.Boolean:
		return NewBoolean(ctx, typename, schema, required)
	case jsonschema.String:
		return NewString(ctx, typename, schema, required)
	case jsonschema.Integer:
		return NewInteger(ctx, typename, schema, required)
	case jsonschema.Number:
		return NewNumber(ctx, typename, schema, required)
	case jsonschema.Array:
		return NewArray(ctx, typename, schema, required)
	case jsonschema.Object:
		fallthrough
	default:
		if RequiesInterface(schema) {
			return NewInterface(ctx, typename, schema, required)
		}

		if schema.Ref != jsonschema.EmptyString {
			def, typename, ctx := ResolvePointer(ctx, schema.Ref)
			return SchemaToType(ctx, typename, def, required)
		}

		return NewAnonymousStruct(ctx, typename, schema, required)
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
