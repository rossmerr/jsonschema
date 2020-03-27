package jsonschema

type Schema struct {
	ID          ID                 `json:"$id,omitempty"`
	Schema      string             `json:"$schema,omitempty"`
	Ref         Reference          `json:"$ref,omitempty"`
	Defs        map[string]*Schema `json:"$defs,omitempty"`
	Anchor      Anchor             `json:"$anchor,omitempty"`
	Description string             `json:"description,omitempty"`
	Title       string             `json:"title,omitempty"`
	Type        Kind               `json:"type,omitempty"`
	Required    []string           `json:"required,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	// Deprecated use Defs
	Definitions          map[string]*Schema `json:"definitions,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	OneOf                []*Schema          `json:"oneof,omitempty"`
	AnyOf                []*Schema          `json:"anyof,omitempty"`
	AllOf                []*Schema          `json:"allof,omitempty"`
	Enum                 []string           `json:"enum,omitempty"`
	AdditionalProperties *bool              `json:"additionalproperties,omitempty"`

	// Validation
	MaxProperties    *uint32 `json:"maxproperties,omitempty"`
	MinProperties    *uint32 `json:"minproperties,omitempty"`
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
}

func (s *Schema) IsEnum() bool {
	return s.Enum != nil
}

func (s *Schema) Structname() string {
	structname := s.ID.Filename()
	if structname == EmptyString {
		structname = s.Type.String()
	}
	return structname
}

func (s *Schema) ArrayType() string {
	arrType := string(s.Items.Type)
	if s.Items.Ref.IsNotEmpty() {
		arrType = s.Items.Ref.Fieldname()
	}
	return arrType
}

func (s *Schema) Stat() (Kind, Reference, []*Schema, []*Schema, []*Schema) {
	return s.Type, s.Ref, s.OneOf, s.AnyOf, s.AllOf
}
