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
	FieldTag   string
	Pointer string
}

func NewNumber(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Number {
	pointer := "*"
	if jsonschema.Contains(required, strings.ToLower(typename)) {
		pointer = ""
	}

	return &Number{
		id:        schema.ID.String(),
		comment:   schema.Description,
		Name:      typename,
		FieldTag:  ctx.Tags.ToFieldTag(typename, schema, required),
		Pointer: pointer,
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
{{ .Name}} {{ .Pointer}}float64 {{ .FieldTag }}
{{- end -}}
`