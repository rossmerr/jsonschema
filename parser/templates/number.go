package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Field = (*Number)(nil)

type Number struct {
	comment   string
	name      string
	fieldTag  string
	Reference string
}

func NewNumber(name, comment string) *Number {
	return &Number{
		name:    name,
		comment: comment,
	}
}

func (s *Number) WithReference(ref bool) parser.Field {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Number) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Number) FieldTag() string {
	return s.fieldTag
}

func (s *Number) Comment() string {
	return s.comment
}

func (s *Number) Name() string {
	return s.name
}

const NumberTemplate = `
{{- define "number" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ typename .Name}} {{ .Reference}}float64 {{ .FieldTag }}
{{- end -}}
`
