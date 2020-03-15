package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Number struct {
	Comment    string
	Name       string
	Validation string
	TypeValue  string
	FieldTag   string
}

func NewNumber(ctx *SchemaContext, key string, schema, parent *jsonschema.Schema) *Number {

	typeValue := schema.Type().String()

	if parent != nil {
		if jsonschema.Contains(parent.Required, key) {
			typeValue = "*" + typeValue
		}
	}

	return &Number{
		Comment:   schema.Description,
		Name:      strings.Title(key),
		TypeValue: typeValue,
		FieldTag:  ctx.Tags.ToFieldTag(key, schema, parent),
	}
}

func (s *Number) types() {}

const NumberTemplate = `
{{- define "number" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .TypeValue}} {{ .FieldTag }}
{{- end -}}
`
