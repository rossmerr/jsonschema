package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
)

type Parser interface {
	Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (map[jsonschema.ID]Component, error)
}

type parser struct {
	packageName string
	registry    *HandlerRegistry
}

func NewParser(packageName string, registry *HandlerRegistry) Parser {
	parser := &parser{
		packageName: packageName,
		registry:    registry,
	}
	return parser
}

func (s *parser) Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (map[jsonschema.ID]Component, error) {
	documents := map[jsonschema.ID]Component{}
	schemaContext := NewSchemaContext(s.packageName, s.Process, references)
	for _, schema := range schemas {
		switch schema.Type {
		case jsonschema.Object:
			t, err := schemaContext.Process(nil, "", schema)
			if err != nil {
				return nil, fmt.Errorf("schemacontext: %w", err)
			}
			documents[schema.ID] = t
		}
	}
	schemaContext.ImplementMethods(documents)
	return documents, nil
}

func (s *parser) Process(schema *jsonschema.Schema, document Root) HandleSchemaFunc {

	var handler HandleSchemaFunc
	switch kind, ref, oneOf, anyOf, allOf, enum := schema.Stat(); {
	case document == nil:
		return s.registry.ResolveHandler(Document)
	case kind == jsonschema.Boolean:
		handler = s.registry.ResolveHandler(Boolean)
	case len(enum) > 0:
		handler = s.registry.ResolveHandler(Enum)
	case kind == jsonschema.String:
		handler = s.registry.ResolveHandler(String)
	case kind == jsonschema.Integer:
		handler = s.registry.ResolveHandler(Interger)
	case kind == jsonschema.Number:
		handler = s.registry.ResolveHandler(Number)
	case kind == jsonschema.Array:
		handler = s.registry.ResolveHandler(Array)
	case ref.IsNotEmpty():
		handler = s.registry.ResolveHandler(Reference)
	case len(oneOf) > 0:
		handler = s.registry.ResolveHandler(OneOf)
	case len(anyOf) > 0:
		handler = s.registry.ResolveHandler(AnyOf)
	case len(allOf) > 0:
		handler = s.registry.ResolveHandler(AllOf)
	default:
		handler = s.registry.ResolveHandler(Object)
	}

	if handler == nil {
		panic("parser: no matching handler was found")
	}

	return handler
}
