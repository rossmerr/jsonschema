package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Boolean)(nil)

type Boolean struct {
	comment   string
	name      string
	FieldTag  string
	Reference string
}

func NewBoolean(name, comment string) *Boolean {
	return &Boolean{
		comment: comment,
		name:    jsonschema.ToTypename(name),
	}
}

func (s *Boolean) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Boolean) WithReference(ref bool) parser.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Boolean) WithFieldTag(tags string) parser.Types {
	s.FieldTag = tags
	return s
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
{{ .Name}} {{ .Reference}}bool {{ .FieldTag }}
{{- end -}}
`
