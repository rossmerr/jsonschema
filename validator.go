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
	case "http://json-schema.org/draft-08/schema#":
		return nil
	case "https://json-schema.org/2019-09/schema":
		return nil
	default:
		return fmt.Errorf("Unknown schema %v, do you need to define a schema flag", schema.Schema)
	}
}
