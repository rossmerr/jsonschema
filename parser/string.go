package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type String struct {
	Comment    string
	Name       string
	Validation string
	FieldTag   string
}

func NewString(ctx *SchemaContext, key string, schema, parent *jsonschema.Schema) *String {
	return &String{
		Comment:  schema.Description,
		Name:     strings.Title(key),
		FieldTag: ctx.Tags.ToFieldTag(key, schema, parent),
	}
}

func (s *String) types() {}

const StringTemplate = `
{{- define "string" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} string {{ .FieldTag }}
{{- end -}}
`
