package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Boolean struct {
	Comment  string
	Name     string
	FieldTag string
}

func NewBoolean(ctx *SchemaContext, key string, schema, parent *jsonschema.Schema) *Boolean {
	return &Boolean{
		Comment:  schema.Description,
		Name:     strings.Title(key),
		FieldTag: ctx.Tags.ToFieldTag(key, schema, parent),
	}
}

func (s *Boolean) types() {}

const BooleanTemplate = `
{{- define "boolean" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} bool
{{- end -}}
`
