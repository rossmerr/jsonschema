package parser

import (
	"context"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
	"github.com/RossMerr/jsonschema/tags/json"
	"github.com/RossMerr/jsonschema/tags/validate"
	"github.com/RossMerr/jsonschema/traversal"
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

	buildReferences(s.ctx, schemas)

	for _, schema := range schemas {
		switch schema.Type {
		case jsonschema.Object:

			anonymousStruct := NewStruct(s.ctx.SetParent(schema), NameFromID(schema.ID), schema.Properties, schema.Description, "", schema.Required...)

			if schema.Defs == nil {
				schema.Defs = map[string]*jsonschema.Schema{}
			}
			for k, v := range schema.Definitions {
				schema.Defs[k] = v
			}

			definitions := make([]Types, 0)
			for typename, def := range schema.Defs {
				definitions = append(definitions, definition(s.ctx.SetParent(schema), NewName(typename), def))
			}
			parse.Structs[schema.ID] = NewDocument(s.ctx, schema.ID.String(), anonymousStruct, definitions, schema.ID.ToFilename())
		}
	}

	return parse
}

func buildReferences(ctx *SchemaContext, schemas map[jsonschema.ID]*jsonschema.Schema) {
	for _, schema := range schemas {
		uris := traversal.WalkSchema(schema)
		for k, v := range uris {
			ctx.References[k] =v
		}
	}
}

func definition(ctx *SchemaContext, name *Name, schema *jsonschema.Schema) *Type {
	t := schemaToType(ctx, name, schema, false)
	arr := ctx.GetMethods(name.Fieldname())
	return PrefixType(t, arr...)
}

func schemaToType(ctx *SchemaContext, name *Name, schema *jsonschema.Schema, renderFieldTags bool, required ...string) Types {
	fieldTag := ""
	if renderFieldTags {
		fieldTag = ctx.Tags.ToFieldTag(name.Tagname(), schema, required)
	}

	isReference := true
	if jsonschema.Contains(required, strings.ToLower(name.Tagname())) {
		isReference = false
	}

	switch kind, ref, oneOf, anyOf, allOf := schema.Stat(); {
	case kind == jsonschema.Boolean:
		return NewBoolean(name, schema.Description, fieldTag, isReference)
	case kind == jsonschema.String:
		if len(schema.Enum) > 0 {
			return NewEnum(ctx, name,schema.Description,fieldTag, isReference,schema.Enum)
		}
		return NewString(name, schema.Description, fieldTag)
	case kind == jsonschema.Integer:
		return NewInteger(name, schema.Description, fieldTag, isReference)
	case kind == jsonschema.Number:
		return NewNumber(name, schema.Description, fieldTag, isReference)
	case kind == jsonschema.Array:
		return NewArray(name, schema.Description, fieldTag, schema.ArrayType())
	case ref.IsNotEmpty():
		return NewReference(ctx, schema.Ref, name, fieldTag)
	case len(oneOf) > 0:
		return NewInterfaceReferenceOneOf(ctx, name, fieldTag, oneOf)
	case len(anyOf) > 0:
		return NewInterfaceReferenceAnyOf(ctx, name, fieldTag, anyOf)
	case len(allOf) > 0:
		return NewInterfaceReferenceAllOf(ctx, name, fieldTag, allOf)
	default:
		return NewStruct(ctx, name, schema.Properties, schema.Description, fieldTag, schema.Required...)
	}
}
