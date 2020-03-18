package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Boolean struct {
	id string
	comment  string
	Name     string
	FieldTag string
}

func NewBoolean(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Boolean {
	return &Boolean{
		id: schema.ID.String(),
		comment:  schema.Description,
		Name:     typename,
		FieldTag: ctx.Tags.ToFieldTag(typename, schema, required),
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
{{ .Name}} bool
{{- end -}}
`
