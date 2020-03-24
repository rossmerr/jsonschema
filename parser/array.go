package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Array struct {
	id string
	comment   string
	Name      string
	TypeValue string
	FieldTag  string
}

func NewArray(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Array {
	arrType := string(schema.Items.Type)
	if schema.Items.Ref != jsonschema.EmptyString {
		arrType = schema.Items.Ref.Fieldname()
	}

	return &Array{
		id: schema.ID.String(),
		comment:   schema.Description,
		Name:     typename,
		TypeValue: arrType,
		FieldTag:  ctx.Tags.ToFieldTag(typename, schema, required),
	}
}

func (s *Array) Comment() string {
	return s.comment
}

func (s *Array) ID() string {
	return s.id
}

const ArrayTemplate = `
{{- define "array" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} []{{ .TypeValue }} {{ .FieldTag }}
{{- end -}}
`
