package types

import "reflect"

const (
	Integer = "integer"
	Number  = "number"
)

type Numeric struct {
	T                string  `json:"Type,omitempty"`
	Description      string  `json:"Description,omitempty"`
	MultipleOf       *uint32 `json:"MultipleOf,omitempty"`
	Maximum          *int32 `json:"Maximum,omitempty"`
	ExclusiveMaximum *int32 `ExclusiveMaximum:"Type,omitempty"`
	Minimum          *int32 `json:"Minimum,omitempty"`
	ExclusiveMinimum *int32 `json:"ExclusiveMinimum,omitempty"`
}

func (s Numeric) Type() reflect.Kind {
	if s.T == Integer {
		return reflect.Int32
	}
	if s.T == Number {
		return reflect.Float64
	}

	return reflect.Invalid
}
