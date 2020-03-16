package parser

import (
	"reflect"

	"github.com/RossMerr/jsonschema"
)

type Array struct {
	id jsonschema.ID
	comment   string
	Name      string
	TypeValue reflect.Kind
	FieldTag  string
}

func NewArray(ctx *SchemaContext, key jsonschema.ID, schema, parent *jsonschema.Schema) *Array {
	return &Array{
		id: key,
		comment:   schema.Description,
		Name:     key.Title(),
		TypeValue: schema.Items.Type(),
		FieldTag:  ctx.Tags.ToFieldTag(key.String(), schema, parent),
	}
}

func (s *Array) Comment() string {
	return s.comment
}

func (s *Array) ID() jsonschema.ID {
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
