package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*Boolean)(nil)
var _ parser.Field = (*Boolean)(nil)

type Boolean struct {
	comment   string
	name      string
	fieldTag  string
	Reference string
}

func NewBoolean(name, comment string) *Boolean {
	return &Boolean{
		comment: comment,
		name:    name,
	}
}



func (s *Boolean) WithReference(ref bool) parser.Field {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Boolean) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Boolean) FieldTag() string {
	return s.fieldTag
}

func (s *Boolean) Comment() string {
	return s.comment
}

func (s *Boolean) Name() string {
	return s.name
}

const BooleanTemplate = `
{{- define "boolean" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ typename .Name}} {{ .Reference}}bool {{ .FieldTag }}
{{- end -}}
`
