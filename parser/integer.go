package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Integer struct {
	id         string
	comment    string
	Name       string
	Validation string
	FieldTag   string
	Pointer string
}

func NewInteger(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Integer {

	pointer := "*"
	if jsonschema.Contains(required, strings.ToLower(typename)) {
		pointer = ""
	}

	return &Integer{
		id:        schema.ID.String(),
		comment:   schema.Description,
		Name:      typename,
		FieldTag:  ctx.Tags.ToFieldTag(typename, schema, required),
		Pointer:pointer,
	}
}

func (s *Integer) Comment() string {
	return s.comment
}

func (s *Integer) ID() string {
	return s.id
}

const IntegerTemplate = `
{{- define "integer" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Pointer}}int32 {{ .FieldTag }}
{{- end -}}
`

