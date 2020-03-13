package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/RossMerr/jsonschema/tokens"
)

const (
	Array = "array"
)

type Arrays struct {
	T           string  `json:"Type,omitempty"`
	Description string  `json:"Description,omitempty"`
	MaxItems    *uint32 `json:"MaxItems,omitempty"`
	MinItems    *uint32 `json:"MinItems,omitempty"`
	UniqueItems bool    `json:"UniqueItems"`
	MaxContains *uint32 `json:"MaxContains,omitempty"`
	MinContains *uint32 `json:"MinContains,omitempty"`
	Items       Type    `json:"-"`
}

func (s Arrays) Type() reflect.Kind {
	if s.T == Array {
		return reflect.Array
	}

	return reflect.Invalid
}

func (s *Arrays) UnmarshalJSON(data []byte) error {

	type Alias Arrays
	err := json.Unmarshal(data, &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	})

	var array map[string]json.RawMessage
	err = json.Unmarshal(data, &array)
	if err != nil {
		return err
	}

	if items, ok := array[tokens.Items]; ok {
		var peek map[string]json.RawMessage
		err = json.Unmarshal(items, &peek)
		if err != nil {
			return err
		}
		if _, ok := peek[tokens.Type]; ok {
			value, err := UnmarshalJSONType(items)
			if err != nil {
				return err
			}
			s.Items = value
		} else if _, ok := peek[tokens.Ref]; ok {
			panic("TODO add support for $Ref")
		} else {
			return fmt.Errorf("Unknown items")
		}
	}

	return err
}
