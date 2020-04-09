package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Field = (*Array)(nil)

type Array struct {
	comment  string
	name     string
	Type     string
	fieldTag string
}

func NewArray(name, comment, arrType string) *Array {
	return &Array{
		comment: comment,
		name:    name,
		Type:    arrType,
	}
}

func (s *Array) WithReference(bool) parser.Field {
	return s
}

func (s *Array) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Array) FieldTag() string {
	return s.fieldTag
}

func (s *Array) Comment() string {
	return s.comment
}

func (s *Array) Name() string {
	return s.name
}

const ArrayTemplate = `
{{- define "array" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ typename .Name}} []{{ .Type }} {{ .FieldTag }}
{{- end -}}
`
