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

	s.findDefinitions(schemas)

	interfaces := s.findInterfaces(schemas)

	for key, schema := range interfaces {
		switch schema.Type() {
		case reflect.Interface:
			inter := NewInterface(s.ctx, key, schema, nil)
			parse.Interfaces[jsonschema.ID(inter.Name)] = inter
		}
	}

	for key, schema := range schemas {
		switch schema.Type() {
		case reflect.Struct:
			parse.Structs[schema.ID] = NewStruct(s.ctx, NewAnonymousStruct(s.ctx, jsonschema.ID(key), schema, nil))
		}
	}

	return parse
}

func (s *parser)findDefinitions(schemas map[jsonschema.ID]*jsonschema.Schema) {
	for key, schema := range schemas {
		s.ctx.Refer[key] = schema
		for key, definition := range schema.Definitions {
			ref := fmt.Sprintf("#/definitions/%v", key)
			s.ctx.Refer[jsonschema.NewID(ref)] = definition
		}
	}
}

func (s *parser)findInterfaces(schemas map[jsonschema.ID]*jsonschema.Schema)  map[jsonschema.ID]*jsonschema.Schema {
	interfaces := map[jsonschema.ID]*jsonschema.Schema{}

	for key, schema := range schemas {
		if 	len(schema.OneOf) > 0 {
			interfaces[key]= schema
		}

		for key, schema = range s.findInterfaces(schema.Properties) {
			interfaces[key] = schema
		}
	}

	return interfaces
}


func SchemaToType(ctx *SchemaContext, key jsonschema.ID, schema, parent *jsonschema.Schema) Types {
	switch schema.Type() {
	case reflect.Struct:
		return NewAnonymousStruct(ctx, jsonschema.ID(key), schema, parent)
	case reflect.Interface:
		return NewInterface(ctx, key, schema, parent)
	case reflect.Array:
		return NewArray(ctx, key, schema, parent)
	case reflect.Int32:
		fallthrough
	case reflect.Float64:
		return NewNumber(ctx, key, schema, parent)
	case reflect.String:
		return NewString(ctx, key, schema, parent)
	case reflect.Ptr:
		ref := ctx.Refer[schema.Ref]
		return SchemaToType(ctx, key, ref, parent)
	default:
		return NewAnonymousStruct(ctx, key, schema, parent)
	}

	return nil
}
