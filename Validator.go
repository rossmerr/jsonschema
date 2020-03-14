package jsonschema

import (
	"fmt"
)

func ValidateSchema(customSchema string, s Schema) error {
	if customSchema != "" {
		if s.Schema == customSchema {
			return nil
		}
	}

	switch s.Schema {
	case "http://json-schema.org/draft-06/schema#":
		return nil
	case "http://json-schema.org/draft-07/schema#":
		return nil
	default:
		return fmt.Errorf("Unknown schema %v, do you need to define a schema flag", s.Schema)
	}
}
