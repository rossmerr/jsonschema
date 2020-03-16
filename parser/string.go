package parser

import (
	"github.com/RossMerr/jsonschema"
)

type String struct {
	id jsonschema.ID
	comment    string
	Name       string
	Validation string
	FieldTag   string
}

func NewString(ctx *SchemaContext, key jsonschema.ID, schema, parent *jsonschema.Schema) *String {
	return &String{
		id: key,
		comment:  schema.Description,
		Name:     key.Title(),
		FieldTag: ctx.Tags.ToFieldTag(key.String(), schema, parent),
	}
}

func (s *String) Comment() string {
	return s.comment
}

func (s *String) ID() jsonschema.ID {
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
