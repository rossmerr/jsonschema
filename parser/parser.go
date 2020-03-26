package parser

import (
	"context"
	"strings"

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
			anonymousStruct := NewStruct(s.ctx, NewName(schema.ID.Filename()), schema, false)

			if schema.Defs == nil {
				schema.Defs = map[string]*jsonschema.Schema{}
			}
			for k, v := range schema.Definitions {
				schema.Defs [k] = v
			}

			definitions := make([]*CustomType, 0)
			for typename, def := range schema.Defs {
				definition := NewDefinition(s.ctx.WrapContext(schema),  NewName(typename), def)
				definitions = append(definitions, definition)
			}
			parse.Structs[schema.ID] = NewDocument(s.ctx, schema.ID.String(), anonymousStruct,definitions, schema.ID.Filename())
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

func SchemaToType(ctx *SchemaContext, name *Name, schema *jsonschema.Schema, renderFieldTags bool, required ...string) Types {
	fieldTag := ""
	if renderFieldTags {
		fieldTag = ctx.Tags.ToFieldTag(name.Tagname(), schema, required)
	}

	isReference := true
	if jsonschema.Contains(required, strings.ToLower(name.Tagname())) {
		isReference = false
	}

	switch schema.Type {
	case jsonschema.Boolean:
		return NewBoolean(name, schema.Description, fieldTag, isReference)
	case jsonschema.String:
		return NewString(name, schema.Description, fieldTag)
	case jsonschema.Integer:
		return NewInteger(name, schema.Description, fieldTag, isReference)
	case jsonschema.Number:
		return NewNumber(name, schema.Description, fieldTag, isReference)
	case jsonschema.Array:
		return NewArray(name, schema.Description, fieldTag, schema.ArrayType())
	case jsonschema.Object:
		fallthrough
	default:
		return NewStruct(ctx, name, schema, renderFieldTags, required...)
	}
}


func NewDefinition(ctx *SchemaContext, name *Name, schema *jsonschema.Schema) *CustomType {
	t := SchemaToType(ctx, name, schema, false)
	arr := ctx.Implementations[name.Fieldname()]
	return PrefixType(t, arr...)
}
