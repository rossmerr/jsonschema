package jsonschema

import (
	"fmt"
)

type Validator interface {
	ValidateSchema(customSchema string, schema Schema) error
}

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (s *validator) ValidateSchema(customSchema string, schema Schema) error {
	if customSchema != EmptyString {
		if schema.Schema == customSchema {
			return nil
		}
	}

	switch schema.Schema {
	case "http://json-schema.org/draft-06/schema#":
		return nil
	case "http://json-schema.org/draft-07/schema#":
		return nil
	default:
		return fmt.Errorf("Unknown schema %v, do you need to define a schema flag", schema.Schema)
	}
}
