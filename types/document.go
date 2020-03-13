package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unicode"

	"github.com/RossMerr/jsonschema/functions"
	"github.com/RossMerr/jsonschema/tokens"
)

const (
	Interface = "interface"
)

type Document struct {
	Properties           map[string]Type `json:"-"`
	Definitions          map[string]Type `json:"-"`
	Description          string          `json:"Description,omitempty"`
	Title                string          `json:"Title,omitempty"`
	Schema               string          `json:"$schema"`
	ID                   string          `json:"$id"`
	T                 string          `json:"type"`
	AdditionalProperties bool            `json:"additionalProperties"`
	Required             []string        `json:"Required"`
}


func (s Document) Type() reflect.Kind {
	if s.T == Interface {
		return reflect.Interface
	}

	return reflect.Struct
}

func (s Document) Filename() string {
	filename := functions.Typename(s.ID)
	return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
}

func (s *Document) UnmarshalJSON(data []byte) error {
	type Alias Document
	err := json.Unmarshal(data, &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	})

	properties, err := PropertiesUnmarshalJSON(data, tokens.Properties)
	if err != nil {
		return err
	}
	s.Properties = properties

	definitions, err := PropertiesUnmarshalJSON(data, tokens.Definitions)
	if err != nil {
		return err
	}
	s.Definitions = definitions

	return err
}

func (s Document) ValidateSchema(customSchema string) error {
	if customSchema != "" {
		if s.Schema == customSchema {
			return nil
		}
	}

	switch s.Schema {
	case "http://json-schema.org/draft-01/schema#":
		return nil
	case "http://json-schema.org/draft-02/schema#":
		return nil
	case "http://json-schema.org/draft-03/schema#":
		return nil
	case "http://json-schema.org/draft-04/schema#":
		return nil
	case "http://json-schema.org/draft-05/schema#":
		return nil
	case "http://json-schema.org/draft-06/schema#":
		return nil
	case "http://json-schema.org/draft-07/schema#":
		return nil
	default:
		return fmt.Errorf("Unknown schema %v, do you need to define a schema flag", s.Schema)
	}
}
