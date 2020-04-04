package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Integer)(nil)

type Integer struct {
	comment   string
	name      string
	fieldTag  string
	Reference string
}

func NewInteger(name, comment string) *Integer {
	return &Integer{
		name:    jsonschema.ToTypename(name),
		comment: comment,
	}
}

func (s *Integer) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Integer) WithReference(ref bool) parser.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Integer) WithFieldTag(tags string) parser.Types {
	s.fieldTag = tags
	return s
}

func (s *Integer) FieldTag() string {
	return s.fieldTag
}

func (s *Integer) Comment() string {
	return s.comment
}

func (s *Integer) Name() string {
	return s.name
}

const IntegerTemplate = `
{{- define "integer" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Reference}}int32 {{ .FieldTag }}
{{- end -}}
`
