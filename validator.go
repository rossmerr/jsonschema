package jsonschema

import (
	"fmt"
)

type Validator interface {
	ValidateSchema(schema Schema) error
}

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (s *validator) ValidateSchema(schema Schema) error {
	switch schema.Schema {
	case "http://json-schema.org/draft-06/schema#":
		return nil
	case "http://json-schema.org/draft-07/schema#":
		return nil
	case "http://json-schema.org/draft-2019-09-16/schema#":
		return nil
	default:
		return fmt.Errorf("Unknown schema %v, do you need to define a schema flag", schema.Schema)
	}
}
