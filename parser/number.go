package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Number struct {
	id         string
	comment    string
	Name       string
	Validation string
	TypeValue  string
	FieldTag   string
}

func NewNumber(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Number {
	typeValue := schema.Type().String()

	if jsonschema.Contains(required, strings.ToLower(typename)) {
		typeValue = "*" + typeValue
	}

	return &Number{
		id:        schema.ID.String(),
		comment:   schema.Description,
		Name:      typename,
		TypeValue: typeValue,
		FieldTag:  ctx.Tags.ToFieldTag(typename, schema, required),
	}
}

func (s *Number) Comment() string {
	return s.comment
}

func (s *Number) ID() string {
	return s.id
}

const NumberTemplate = `
{{- define "number" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .TypeValue}} {{ .FieldTag }}
{{- end -}}
`
