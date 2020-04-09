package jsonschema

import "encoding/json"

// Schema the base JSON Schema
type Schema struct {
	// Annotations
	Description string `json:"description,omitempty"`
	Title       string `json:"title,omitempty"`
	Default     string `json:"default,omitempty"`
	Examples    string `json:"examples,omitempty"`

	ID     ID            `json:"$id,omitempty"`
	Schema SchemaVersion `json:"$schema,omitempty"`
	Ref    Reference     `json:"$ref,omitempty"`

	Defs map[string]*Schema `json:"$defs,omitempty"`
	// Deprecated use Defs
	Definitions map[string]*Schema `json:"definitions,omitempty"`

	Anchor string `json:"$anchor,omitempty"`

	Type DataType `json:"type,omitempty"`

	// Required Properties
	Required []string `json:"required,omitempty"`

	// Properties
	Properties           map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties *bool              `json:"additionalproperties,omitempty"`

	Items *Schema   `json:"items,omitempty"`
	OneOf []*Schema `json:"oneof,omitempty"`
	AnyOf []*Schema `json:"anyof,omitempty"`
	AllOf []*Schema `json:"allof,omitempty"`
	Enum  []string  `json:"enum,omitempty"`

	// Size
	MaxProperties *uint32 `json:"maxproperties,omitempty"`
	MinProperties *uint32 `json:"minproperties,omitempty"`

	// Validation
	MaxLength        *uint32 `json:"maxlength,omitempty"`
	MinLength        *uint32 `json:"minlength,omitempty"`
	MaxContains      *uint32 `json:"maxcontains,omitempty"`
	MinContains      *uint32 `json:"mincontains,omitempty"`
	MaxItems         *uint32 `json:"maxitems,omitempty"`
	MinItems         *uint32 `json:"minitems,omitempty"`
	Maximum          *int32  `json:"maximum,omitempty"`
	ExclusiveMaximum *int32  `json:"exclusivemaximum,omitempty"`
	Minimum          *int32  `json:"minimum,omitempty"`
	ExclusiveMinimum *int32  `json:"exclusiveminimum,omitempty"`
	Pattern          string  `json:"pattern,omitempty"`

	Parent *Schema `json:"-"`
	Key    string  `json:"-"`
}

// SetParent recursively traverse the schema setting any parents and keys
func (s *Schema) SetParent(key string, parent *Schema) {
	s.Parent = parent
	s.Key = key
	for k, subschema := range s.Properties {
		subschema.SetParent(k, s)
	}
	for k, subschema := range s.Defs {
		subschema.SetParent(k, s)
	}
	for k, subschema := range s.Definitions {
		subschema.SetParent(k, s)
	}
	for _, subschema := range s.AllOf {
		subschema.SetParent(key, s)
	}
	for _, subschema := range s.AnyOf {
		subschema.SetParent(key, s)
	}
	for _, subschema := range s.OneOf {
		subschema.SetParent(key, s)
	}
}

// Base finds the root/base schema from any point in the schema hierarchy
func (s *Schema) Base() *Schema {
	return root(s)
}

func root(s *Schema) *Schema {
	if s.Parent == nil {
		return s
	}
	return root(s.Parent)
}

func (s *Schema) UnmarshalJSON(b []byte) error {
	type Alias Schema
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	return nil
}

// AllDefinitions returns a merged map of the definitions and $def fields, $def takes precedences
func (s *Schema) AllDefinitions() map[string]*Schema {
	definitions := map[string]*Schema{}
	for key, def := range s.Definitions {
		definitions[key] = def
	}

	for key, def := range s.Defs {
		definitions[key] = def
	}
	return definitions
}

// Stat returns the main keywords of the schema to workout how to process it
func (s *Schema) Stat() (DataType, Reference, []*Schema, []*Schema, []*Schema, []string) {
	return s.Type, s.Ref, s.OneOf, s.AnyOf, s.AllOf, s.Enum
}
