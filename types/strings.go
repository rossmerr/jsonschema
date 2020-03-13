package types

import "reflect"

const (
	String = "string"
)

type Strings struct {
	T           string  `json:"Type,omitempty"`
	Description string  `json:"Description,omitempty"`
	MaxLength   *uint32 `json:"MaxLength,omitempty"`
	MinLength   *uint32 `json:"MinLength,omitempty"`
	Pattern     string  `json:"Pattern,omitempty"`
}

func (s Strings) Type() reflect.Kind {
	if s.T == String {
		return reflect.String
	}

	return reflect.Invalid
}
