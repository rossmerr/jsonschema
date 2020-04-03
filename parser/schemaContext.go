package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal/traverse"
)

type Resolve func(name string, schema *jsonschema.Schema) HandleSchemaFunc

// NewSchemaContext return's a SchemaContext
func NewSchemaContext(resolve Resolve, references map[jsonschema.ID]*jsonschema.Schema) *SchemaContext {
	return &SchemaContext{
		implementations: map[string][]*Method{},
		documents:       map[string]*Document{},
		resolve:         resolve,
		references:      references,
	}
}

// SchemaContext is the collection of all schema's being processed
type SchemaContext struct {
	implementations map[string][]*Method
	resolve         Resolve
	references      map[jsonschema.ID]*jsonschema.Schema
	documents       map[string]*Document
	document        *Document
}

// NewDocument return's and process a Document from a schema
func (s *SchemaContext) NewDocument(id, packageName, filename string, schema *jsonschema.Schema) (*Document, error) {
	if schema.Key == "" {
		schema.SetParent(schema.ID.ToTypename(), nil)
	}

	s.document = &Document{
		ID:         id,
		Package:    packageName,
		Globals:    map[string]Types{},
		Filename:   filename,
		rootSchema: schema,
	}

	t, err := s.Process(schema.ID.ToTypename(), schema)
	if err != nil {
		return nil, fmt.Errorf("schemacontext: %w", err)
	}

	s.document.Globals[t.Name()] = t
	s.documents[id] = s.document
	return s.document, nil
}

// Dispose add's any methods that any struct might need to implement for any interfaces
func (s *SchemaContext) Dispose() {
	for _, doc := range s.documents {
		for k, g := range doc.Globals {
			arr := s.implementations[k]
			if len(arr) > 0 {
				g.WithMethods(arr...)
			}
		}
	}
}

// ImplementInterface add's any methods onto the named receiver across the entire schema
// so you can implement a interface from a reference etc
func (s *SchemaContext) ImplementInterface(receiver string, methods ...*Method) {
	if receiver != jsonschema.EmptyString {
		switch arr, ok := s.implementations[receiver]; {
		case !ok:
			arr = []*Method{}
			fallthrough
		default:
			arr = append(arr, methods...)
			s.implementations[receiver] = arr
		}
	}
}

// Process a schema and return it as a tree of Types
func (s *SchemaContext) Process(name string, schema *jsonschema.Schema) (Types, error) {
	if s.document == nil {
		panic(fmt.Errorf("schemacontext: document not set %v", s.document))
	}
	handler := s.resolve(name, schema)
	return handler(s, s.document, name, schema)
}

// ResolvePointer takes a Reference and uses it to walk the schema to find any types to reference
func (s *SchemaContext) ResolvePointer(ref jsonschema.Reference) (string, error) {
	path := ref.Path()
	if fieldname := path.ToFieldname(); fieldname != "." {
		return fieldname, nil
	}

	var base *jsonschema.Schema
	base = s.document.Root()
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
	return def.ID.ToTypename(), nil
}
