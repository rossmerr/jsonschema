package jsonschema

import (
	"reflect"
)

type Schema struct {
	ID                   ID                 `json:"$id,omitempty"`
	Schema               string             `json:"$schema,omitempty"`
	Ref                  ID             `json:"$ref,omitempty"`
	Description          string             `json:"Description,omitempty"`
	Title                string             `json:"Title,omitempty"`
	TypeValue            string             `json:"Type,omitempty"`
	Required             []string           `json:"Required,omitempty"`
	Properties           map[ID]*Schema `json:"Properties,omitempty"`
	Definitions          map[ID]*Schema `json:"Definitions,omitempty"`
	AdditionalProperties bool               `json:"additionalProperties"`
	Items                *Schema            `json:"Items,omitempty"`
	OneOf                []*Schema          `json:"oneOf,omitempty"`

	// Validation
	MaxProperties    *uint32 `json:"MaxProperties,omitempty"`
	MinProperties    *uint32 `json:"MinProperties,omitempty"`
	MaxLength        *uint32 `json:"MaxLength,omitempty"`
	MinLength        *uint32 `json:"MinLength,omitempty"`
	MaxContains      *uint32 `json:"MaxContains,omitempty"`
	MinContains      *uint32 `json:"MinContains,omitempty"`
	MaxItems         *uint32 `json:"MaxItems,omitempty"`
	MinItems         *uint32 `json:"MinItems,omitempty"`
	Maximum          *int32  `json:"Maximum,omitempty"`
	ExclusiveMaximum *int32  `ExclusiveMaximum:"Type,omitempty"`
	Minimum          *int32  `json:"Minimum,omitempty"`
	ExclusiveMinimum *int32  `json:"ExclusiveMinimum,omitempty"`
}

func (s Schema) Type() reflect.Kind {
	switch s.TypeValue {
	case Number:
		return reflect.Float64
	case Integer:
		return reflect.Int32
	case String:
		return reflect.String
	case Array:
		return reflect.Array
	case Object:
		if s.OneOf != nil && len(s.OneOf) > 0 {
			return reflect.Interface
		}
		if s.Ref != EmptyString {
			return reflect.Ptr
		}
		return reflect.Struct
	}

	if s.Ref != EmptyString {
		return reflect.Ptr
	}

	return reflect.Invalid
}
