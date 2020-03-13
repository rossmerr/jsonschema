package types

import (
	"encoding/json"
	"fmt"

	"github.com/RossMerr/jsonschema/tokens"
)

func PropertiesUnmarshalJSON(data []byte, token string) (map[string]Type, error) {
	s := map[string]Type{}

	var array map[string]json.RawMessage
	err := json.Unmarshal(data, &array)
	if err != nil {
		return s, err
	}

	if items, ok := array[token]; ok {

		var properties map[string]json.RawMessage
		err = json.Unmarshal(items, &properties)
		if err != nil {
			return s, err
		}

		for key, propertie := range properties {

			value, err := UnmarshalJSONType(propertie)
			if err != nil {
				return s, err
			}

			s[key] = value
		}
	}
	return s, nil
}

func UnmarshalJSONType(data json.RawMessage) (Type, error) {
	var peek map[string]json.RawMessage
	err := json.Unmarshal(data, &peek)
	if err != nil {
		return nil, err
	}

	if t, ok := peek[tokens.Type]; ok {
		var tt string
		err := json.Unmarshal(t, &tt)
		if err != nil {
			return nil, err
		}

		switch Types(tt) {
		case Array:
			var value Arrays
			err := value.UnmarshalJSON(data)
			// TODO
			//err := json.Unmarshal(data, &value)
			return value, err
		case String:
			var value Strings
			err := json.Unmarshal(data, &value)
			return value, err

		case Integer:
			var value Numeric
			err := json.Unmarshal(data, &value)
			return value, err
		case Number:
			var value Numeric
			err := json.Unmarshal(data, &value)
			return value, err
		default:
			var value Objects
			err := json.Unmarshal(data, &value)
			return value, err
		}
		return nil, fmt.Errorf("Unknown type %v", t)
	}

	if _, ok := peek[tokens.Ref]; ok {
		var value References
		err := json.Unmarshal(data, &value)
		return value, err
	}

	var value Objects
	err = json.Unmarshal(data, &value)
	return value, err
}
