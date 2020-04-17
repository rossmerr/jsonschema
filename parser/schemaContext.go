package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal/traverse"
)

// NewSchemaContext return's a SchemaContext
func NewSchemaContext(packageName string, resolve Resolve, references map[jsonschema.ID]*jsonschema.Schema) *SchemaContext {
	return &SchemaContext{
		implementations: map[string][]*MethodSignature{},
		resolve:         resolve,
		references:      references,
		PackageName:     packageName,
	}
}

// SchemaContext is the collection of all schema's being processed
type SchemaContext struct {
	implementations map[string][]*MethodSignature
	resolve         Resolve
	references      map[jsonschema.ID]*jsonschema.Schema
	document        Root
	PackageName     string
}

// ImplementMethods add's any methods that any struct might need to implement for any interfaces
func (s *SchemaContext) ImplementMethods(documents map[jsonschema.ID]Component) {
	for _, doc := range documents {
		if document, ok := doc.(Root); ok {
			for k, g := range document.Globals() {
				obj, ok := g.(Receiver)
				if !ok {
					continue
				}
				methodSignatures := s.implementations[k]
				for _, methodSignature := range methodSignatures {
					method := NewMethodFromSignature(k, methodSignature)
					obj.WithMethods(method)
				}
			}
		}
	}
}

// RegisterMethodSignature add's any methods onto the named receiver across all schemas
// so you can implement a interface from a reference etc
func (s *SchemaContext) RegisterMethodSignature(receiver string, methods ...*MethodSignature) {
	if receiver != emptyString {
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
func (s *SchemaContext) Process(document Root, name string, schema *jsonschema.Schema) (Component, error) {
	handler := s.resolve(schema, document)
	return handler(s, document, name, schema)
}

// ResolveSchema takes a Reference and uses it to walk the schema to find the matching subschema
func (s *SchemaContext) ResolveSchema(ref jsonschema.Reference, arr ...*jsonschema.Schema) (*jsonschema.Schema, error) {
	path := ref.Path()

	var base *jsonschema.Schema
	if len(arr) > 0 {
		base = arr[0]
	}

	if id, err := ref.ID(); err == nil {
		base = s.references[id]
	}

	if base == nil {
		return nil, fmt.Errorf("resolvepointer: base schema not provided or not found")
	}

	def := traverse.Walk(base, path)
	if def == nil {
		return nil, fmt.Errorf("resolvepointer: path not found %v", path)
	}

	return def, nil
}

// ResolveTypename takes a Reference and uses it to walk the schema to find the matching typename to reference
func (s *SchemaContext) ResolveTypename(ref jsonschema.Reference, base *jsonschema.Schema) (string, error) {
	path := ref.Path()
	if fieldname := path.ToKey(); fieldname != "." {
		return fieldname, nil
	}

	refSchema, err := s.ResolveSchema(ref, base)
	if err != nil {
		return ".", err
	}

	return refSchema.ID.Fragment(), nil
}
