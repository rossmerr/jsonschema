package jsonschema

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Schema struct {
	ID          ID                 `json:"$id,omitempty"`
	Schema      string             `json:"$schema,omitempty"`
	Ref         Pointer            `json:"$ref,omitempty"`
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
	AllOf                map[string]*Schema          `json:"allof,omitempty"`
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

	// All unhandled json fields are unmarshaled here
	UnknownFields map[string]interface{} `json:"-"`
}

func (s *Schema) IsEnum() bool {
	return s.Enum != nil
}

func (s *Schema) UnmarshalJSON(b []byte) (err error) {

	type Alias Schema
	a := Alias{}

	if err = json.Unmarshal(b, &a); err == nil {
		*s = Schema(a)
		s.UnknownFields = make(map[string]interface{})

	}

	if err != nil {
		return
	}
	m := make(map[string]interface{})

	if err = json.Unmarshal(b, &m); err == nil {

		for k, v := range m {
			delete(m, k)
			m[strings.ToLower(k)] = v
		}

		val := reflect.ValueOf(s).Elem()
		for i := 0; i < val.NumField(); i++ {
			tag := val.Type().Field(i).Tag
			if v, ok := tag.Lookup("json"); ok {
				tagFields := strings.Split(v, ",")
				ForEach(tagFields, func(v string) string {
					delete(m, v)
					return v
				})
			}
		}

		for k, v := range m {
			s.UnknownFields[k] = v
		}
	}

	return
}
