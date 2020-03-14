package jsonschema

import (
	"reflect"
)

const (
	Object = "object"
	Array = "array"
	Integer = "integer"
	Number  = "number"
	String = "string"
)

type Schema struct {
	ID                   string          `json:"$id,omitempty"`
	Schema               string          `json:"$schema,omitempty"`
	Description          string          `json:"Description,omitempty"`
	Title                string          `json:"Title,omitempty"`
	AdditionalProperties bool            `json:"additionalProperties"`
	TypeValue             string          `json:"Type,omitempty"`
	MaxProperties *uint32         `json:"MaxProperties,omitempty"`
	MinProperties *uint32         `json:"MinProperties,omitempty"`
	Required      []string        `json:"Required,omitempty"`
	OneOf         []*Schema           `json:"oneOf"`
	Properties    map[string]*Schema `json:"Properties,omitempty"`
	Definitions          map[string]*Schema `json:"-"`
	Items       *Schema    `json:"Items,omitempty"`


	// Validation
	MaxLength   *uint32 `json:"MaxLength,omitempty"`
	MinLength   *uint32 `json:"MinLength,omitempty"`
	MaxContains *uint32 `json:"MaxContains,omitempty"`
	MinContains *uint32 `json:"MinContains,omitempty"`
	MaxItems    *uint32 `json:"MaxItems,omitempty"`
	MinItems    *uint32 `json:"MinItems,omitempty"`
	Maximum          *int32 `json:"Maximum,omitempty"`
	ExclusiveMaximum *int32 `ExclusiveMaximum:"Type,omitempty"`
	Minimum          *int32 `json:"Minimum,omitempty"`
	ExclusiveMinimum *int32 `json:"ExclusiveMinimum,omitempty"`
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
		return reflect.Struct
	}

	return reflect.Invalid
}