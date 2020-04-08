package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal/traverse"
)

type Resolve func(name string, schema *jsonschema.Schema, document *Document) HandleSchemaFunc

// NewSchemaContext return's a SchemaContext
func NewSchemaContext(packageName string, resolve Resolve, references map[jsonschema.ID]*jsonschema.Schema) *SchemaContext {
	return &SchemaContext{
		implementations: map[string][]*MethodSignature{},
		resolve:         resolve,
		references:      references,
		packageName:packageName,
	}
}

// SchemaContext is the collection of all schema's being processed
type SchemaContext struct {
	implementations map[string][]*MethodSignature
	resolve         Resolve
	references      map[jsonschema.ID]*jsonschema.Schema
	document        *Document
	packageName string
}

// ImplementMethods add's any methods that any struct might need to implement for any interfaces
func (s *SchemaContext) ImplementMethods(documents map[jsonschema.ID]Types) {
	for _, doc := range documents {
		if document, ok  := doc.(*Document); ok {
			document.WithPackageName(s.packageName)
			for k, g := range document.Globals {
				methodSignatures := s.implementations[k]
				for _, methodSignature := range methodSignatures {
					method := NewMethodFromSignature(k, methodSignature)
					g.WithMethods(method)
				}
			}
		}
	}
}

// RegisterMethodSignature add's any methods onto the named receiver across all schemas
// so you can implement a interface from a reference etc
func (s *SchemaContext) RegisterMethodSignature(receiver string, methods ...*MethodSignature) {
	if receiver != EmptyString {
		switch arr, ok := s.implementations[receiver]; {
		case !ok:
			arr = []*MethodSignature{}
			fallthrough
		default:
			arr = append(arr, methods...)
			s.implementations[receiver] = arr
		}
	}
}

// Process a schema and return it as a tree of Types
func (s *SchemaContext) Process(document *Document, name string, schema *jsonschema.Schema) (Types, error) {
	handler := s.resolve(name, schema, document)
	return handler(s, document, name, schema)
}

// ResolvePointer takes a Reference and uses it to walk the schema to find any types to reference
func (s *SchemaContext) ResolvePointer(ref jsonschema.Reference, doc *Document) (string, error) {
	path := ref.Path()
	if fieldname := path.ToKey(); fieldname != "." {
		return fieldname, nil
	}

	var base *jsonschema.Schema
	base = doc.Root()
	if id, err := ref.ID(); err == nil {
		if err != nil {
			return ".", fmt.Errorf("resolvepointer: %w", err)

		}
		base = s.references[id]
	}

	def := traverse.Walk(base, path)
	if def == nil {
		return ".", fmt.Errorf("resolvepointer: path not found %v", path)
	}

	return def.ID.ToKey(), nil
}
