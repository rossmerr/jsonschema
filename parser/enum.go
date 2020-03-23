package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Enum struct {
	id       string
	comment  string
	Name     string
	Values   []string
	FieldTag string
	Pointer  string
}

func NewEnum(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Enum {
	pointer := "*"
	if jsonschema.Contains(required, strings.ToLower(typename)) {
		pointer = ""
	}

	return &Enum{
		id:       schema.ID.String(),
		comment:  schema.Description,
		Name:     typename,
		FieldTag: ctx.Tags.ToFieldTag(typename, schema, required),
		Pointer:  pointer,
		Values:   schema.Enum,
	}
}

func (s *Enum) Comment() string {
	return s.comment
}

func (s *Enum) ID() string {
	return s.id
}

const EnumTemplate = `
{{- define "enum" -}}


// test
{{end -}}
`
