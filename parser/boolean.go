package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Boolean struct {
	id jsonschema.ID
	comment  string
	Name     string
	FieldTag string
}

func NewBoolean(ctx *SchemaContext, key jsonschema.ID, schema, parent *jsonschema.Schema) *Boolean {
	return &Boolean{
		id: key,
		comment:  schema.Description,
		Name:     key.Title(),
		FieldTag: ctx.Tags.ToFieldTag(key.String(), schema, parent),
	}
}

func (s *Boolean) Comment() string {
	return s.comment
}

func (s *Boolean) ID() jsonschema.ID {
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
