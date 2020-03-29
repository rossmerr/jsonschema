package jsonschema

import (
	"fmt"
	"strings"
)

var schemas = []string{"http://json-schema.org/draft-07/schema", "http://json-schema.org/draft-08/schema", "https://json-schema.org/2019-09/schema"}

type Validator interface {
	ValidateSchema(id ID, schema *Schema) error
}

type validator struct {
}

func NewValidator() Validator {
	return &validator{}
}

func (s *validator) ValidateSchema(id ID, schema *Schema) error {
	trim := strings.TrimRight(schema.Schema, "#")
	if !Contains(schemas, trim) {
		return fmt.Errorf("Unsupported schema found %v in %v\n\nTry using one of:\n%v\n", schema.Schema, id, strings.Join(schemas, ", "))
	}
	return nil
}
