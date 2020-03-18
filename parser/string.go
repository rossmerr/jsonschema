package parser

import (
	"github.com/RossMerr/jsonschema"
)

type String struct {
	id string
	comment    string
	Name       string
	Validation string
	FieldTag   string
}

func NewString(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *String {
	return &String{
		id: schema.ID.String(),
		comment:  schema.Description,
		Name:     typename,
		FieldTag: ctx.Tags.ToFieldTag(typename, schema, required),
	}
}

func (s *String) Comment() string {
	return s.comment
}

func (s *String) ID() string {
	return s.id
}

const StringTemplate = `
{{- define "string" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} string {{ .FieldTag }}
{{- end -}}
`
