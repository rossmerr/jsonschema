package types

import "reflect"

type References struct {
	Ref string`json:"$ref,omitempty"`
	Description string `json:"Description,omitempty"`
}

func (s References) Type() reflect.Kind {
	return reflect.Ptr
}