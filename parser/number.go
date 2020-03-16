package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Number struct {
	id jsonschema.ID
	comment    string
	Name       string
	Validation string
	TypeValue  string
	FieldTag   string
}

func NewNumber(ctx *SchemaContext, key jsonschema.ID, schema, parent *jsonschema.Schema) *Number {

	typeValue := schema.Type().String()

	if parent != nil {
		if jsonschema.Contains(parent.Required, key.String()) {
			typeValue = "*" + typeValue
		}
	}

	return &Number{
		id: key,
		comment:   schema.Description,
		Name:      key.Title(),
		TypeValue: typeValue,
		FieldTag:  ctx.Tags.ToFieldTag(key.String(), schema, parent),
	}
}

func (s *Number) Comment() string {
	return s.comment
}

func (s *Number) ID() jsonschema.ID {
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
