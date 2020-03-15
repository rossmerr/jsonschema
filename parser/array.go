package parser

import (
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Array struct {
	Comment   string
	Name      string
	TypeValue reflect.Kind
	FieldTag  string
}

func NewArray(ctx *SchemaContext, key string, schema, parent *jsonschema.Schema) *Array {
	return &Array{
		Comment:   schema.Description,
		Name:      strings.Title(key),
		TypeValue: schema.Items.Type(),
		FieldTag:  ctx.Tags.ToFieldTag(key, schema, parent),
	}
}

func (s *Array) types() {}

const ArrayTemplate = `
{{- define "array" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} []{{ .TypeValue }} {{ .FieldTag }}
{{- end -}}
`
