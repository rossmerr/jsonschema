package parser

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/tags"
	"github.com/RossMerr/jsonschema/parser/tags/json"
	"github.com/RossMerr/jsonschema/parser/tags/validate"
)

type Parser interface {
	Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (*Parse, error)
}

type parser struct {
	ctx *SchemaContext
}

func NewParser(ctx context.Context, packageName string) Parser {
	return &parser{
		ctx: NewContext(ctx, packageName, tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})),
	}
}

func (s *parser) Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (*Parse, error) {
	parse := NewParse()

	s.ctx.References = references

	for _, schema := range schemas {
		switch schema.Type {
		case jsonschema.Object:

			anonymousStruct, err := NewStruct(s.ctx.SetParent(schema), NameFromID(schema.ID), schema.Properties, schema.Description, "", schema.Required...)
			if err != nil {
				return nil, err
			}

			if schema.Defs == nil {
				schema.Defs = map[string]*jsonschema.Schema{}
			}
			for k, v := range schema.Definitions {
				schema.Defs[k] = v
			}

			for typename, def := range schema.Defs {
				d, err := definition(s.ctx.SetParent(schema), NewName(typename), def)
				if err != nil {
					return nil, err
				}
				if _, ok := s.ctx.Globals[typename]; !ok {
					s.ctx.Globals[typename] = d
				} else {
					return nil, fmt.Errorf("Global keys need to be unique found %v more than once, last references was in %v", typename, schema.ID)
				}
			}

			parse.Structs[schema.ID] = NewDocument(s.ctx, schema.ID.String(), anonymousStruct, toFilename(schema.ID))
		}
	}

	return parse, nil
}

// toFilename returns the file name from the ID.
func toFilename(s jsonschema.ID) string {
	name := s.ToTypename()

	if len(name) > 0 {
		return string(unicode.ToLower(rune(name[0]))) + name[1:]
	}
	return name
}

func definition(ctx *SchemaContext, name *Name, schema *jsonschema.Schema) (*Type, error) {
	t, err := schemaToType(ctx, name, schema, false)
	if err != nil {
		return nil, err
	}
	arr := ctx.GetMethods(name.Fieldname())
	return PrefixType(t, arr...), nil
}

func schemaToType(ctx *SchemaContext, name *Name, schema *jsonschema.Schema, renderFieldTags bool, required ...string) (Types, error) {
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
		return NewBoolean(name, schema.Description, fieldTag, isReference), nil
	case kind == jsonschema.String:
		if len(schema.Enum) > 0 {
			return NewEnum(ctx, name, schema.Description, fieldTag, isReference, schema.Enum)
		}
		return NewString(name, schema.Description, fieldTag), nil
	case kind == jsonschema.Integer:
		return NewInteger(name, schema.Description, fieldTag, isReference), nil
	case kind == jsonschema.Number:
		return NewNumber(name, schema.Description, fieldTag, isReference), nil
	case kind == jsonschema.Array:
		return NewArray(name, schema.Description, fieldTag, ArrayType(schema)), nil
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

func ArrayType(s *jsonschema.Schema) string {
	arrType := string(s.Items.Type)
	if s.Items.Ref.IsNotEmpty() {
		arrType = s.Items.Ref.ToTypename()
	}
	return arrType
}
