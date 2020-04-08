package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*Integer)(nil)
var _ parser.Field = (*Integer)(nil)

type Integer struct {
	comment   string
	name      string
	fieldTag  string
	Reference string
}

func NewInteger(name, comment string) *Integer {
	return &Integer{
		name:    name,
		comment: comment,
	}
}

func (s *Integer) WithReference(ref bool) parser.Field {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Integer) WithFieldTag(tags string) parser.Field {
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
{{ typename .Name}} {{ .Reference}}int32 {{ .FieldTag }}
{{- end -}}
`
