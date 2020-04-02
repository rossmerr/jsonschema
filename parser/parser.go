package parser

import (
	"fmt"
	"unicode"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

type Parser interface {
	Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (map[jsonschema.ID]*document.Document, error)
	HandlerFunc(kind Kind, handler document.HandleSchemaFunc)
}

type parser struct {
	handlers    map[Kind]document.HandleSchemaFunc
	packageName string
}

func NewParser(packageName string) Parser {
	parser := &parser{
		packageName: packageName,
		handlers:    map[Kind]document.HandleSchemaFunc{},
	}
	return parser
}

func (s *parser) Parse(schemas map[jsonschema.ID]*jsonschema.Schema, references map[jsonschema.ID]*jsonschema.Schema) (map[jsonschema.ID]*document.Document, error) {
	documents := map[jsonschema.ID]*document.Document{}

	for _, schema := range schemas {
		switch schema.Type {
		case jsonschema.Object:
			doc := document.NewDocument(schema.ID.String(), s.packageName, toFilename(schema.ID), schema, s.Process, references)

			s, err := doc.Process(schema.ID.ToTypename(), schema)
			if err != nil {
				return nil, fmt.Errorf("parse: %w", err)
			}

			doc.WithType(s)

			documents[schema.ID] = doc
		}
	}

	return documents, nil
}

func (s *parser) HandlerFunc(kind Kind, handler document.HandleSchemaFunc) {
	if _, ok := s.handlers[kind]; ok {
		panic(fmt.Sprintf("parser: multiple registrations for %v", kind))
	} else {
		s.handlers[kind] = handler
	}
}

func (s *parser) Process(name string, schema *jsonschema.Schema) document.HandleSchemaFunc {
	var handler document.HandleSchemaFunc
	switch kind, ref, oneOf, anyOf, allOf, enum, isParent := schema.Stat(); {
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
	case isParent:
		handler = s.handlers[RootObject]
	default:
		handler = s.handlers[Object]
	}

	if handler == nil {
		panic("parser: no matching handler was found")
	}

	return handler
}

// toFilename returns the file name from the ID.
func toFilename(s jsonschema.ID) string {
	name := s.ToTypename()

	if len(name) > 0 {
		return string(unicode.ToLower(rune(name[0]))) + name[1:]
	}
	return name
}
