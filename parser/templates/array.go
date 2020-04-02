package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Array)(nil)

type Array struct {
	comment  string
	name     string
	Type     string
	FieldTag string
}

func NewArray(name, comment, arrType string) *Array {
	return &Array{
		comment: comment,
		name:    jsonschema.ToTypename(name),
		Type:    arrType,
	}

}

func (s *Array)	WithMethods(methods ...string) parser.Types {
	return s
}

func (s *Array) WithReference(ref bool) parser.Types {
	return s
}

func (s *Array) WithFieldTag(tags string) parser.Types {
	s.FieldTag = tags
	return s
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
{{ .Name}} []{{ .Type }} {{ .FieldTag }}
{{- end -}}
`
