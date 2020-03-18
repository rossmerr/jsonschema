package parser

import (
	"reflect"

	"github.com/RossMerr/jsonschema"
)

type Array struct {
	id string
	comment   string
	Name      string
	TypeValue reflect.Kind
	FieldTag  string
}

func NewArray(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Array {
	return &Array{
		id: schema.ID.String(),
		comment:   schema.Description,
		Name:     typename,
		TypeValue: schema.Items.Type(),
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
