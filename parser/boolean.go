package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Boolean struct {
	id string
	comment  string
	Name     string
	FieldTag string
	Pointer string

}

func NewBoolean(ctx *SchemaContext, field string, schema *jsonschema.Schema, required []string) *Boolean {
	pointer := "*"
	if jsonschema.Contains(required, strings.ToLower(field)) {
		pointer = ""
	}

	return &Boolean{
		id:       schema.ID.String(),
		comment:  schema.Description,
		Name:     field,
		FieldTag: ctx.Tags.ToFieldTag(field, schema, required),
		Pointer:  pointer,
	}
}

func (s *Boolean) Comment() string {
	return s.comment
}

func (s *Boolean) ID() string {
	return s.id
}

const BooleanTemplate = `
{{- define "boolean" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Pointer}}bool {{ .FieldTag }}
{{- end -}}
`
