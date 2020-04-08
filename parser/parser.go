package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
)

type Parser interface {
	Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (map[jsonschema.ID]Types, error)
	HandlerFunc(kind Kind, handler HandleSchemaFunc)
}

type parser struct {
	handlers    map[Kind]HandleSchemaFunc
	packageName string
}

func NewParser(packageName string) Parser {
	parser := &parser{
		packageName: packageName,
		handlers:    map[Kind]HandleSchemaFunc{},
	}
	return parser
}

func (s *parser) Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (map[jsonschema.ID]Types, error) {
	documents := map[jsonschema.ID]Types{}
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

func (s *parser) HandlerFunc(kind Kind, handler HandleSchemaFunc) {
	if _, ok := s.handlers[kind]; ok {
		panic(fmt.Sprintf("parser: multiple registrations for %v", kind))
	} else {
		s.handlers[kind] = handler
	}
}

func (s *parser) Process(name string, schema *jsonschema.Schema, document *Document) HandleSchemaFunc {

	var handler HandleSchemaFunc
	switch kind, ref, oneOf, anyOf, allOf, enum := schema.Stat(); {
	case document == nil:
		return s.handlers[RootDocument]
	case kind == jsonschema.Boolean:
		handler = s.handlers[Boolean]
	case len(enum) > 0:
		handler = s.handlers[Enum]
	case kind == jsonschema.String:
		handler = s.handlers[String]
	case kind == jsonschema.Integer:
		handler = s.handlers[Interger]
	case kind == jsonschema.Number:
		handler = s.handlers[Number]
	case kind == jsonschema.Array:
		handler = s.handlers[Array]
	case ref.IsNotEmpty():
		handler = s.handlers[Reference]
	case len(oneOf) > 0:
		handler = s.handlers[OneOf]
	case len(anyOf) > 0:
		handler = s.handlers[AnyOf]
	case len(allOf) > 0:
		handler = s.handlers[AllOf]
	default:
		handler = s.handlers[Object]
	}

	if handler == nil {
		panic("parser: no matching handler was found")
	}

	return handler
}
