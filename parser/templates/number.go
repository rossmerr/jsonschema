package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Number)(nil)

type Number struct {
	comment   string
	name      string
	FieldTag  string
	Reference string
}

func NewNumber(name, comment string) *Number {
	return &Number{
		name:    jsonschema.ToTypename(name),
		comment: comment,
	}
}

func (s *Number) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Number) WithReference(ref bool) parser.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Number) WithFieldTag(tags string) parser.Types {
	s.FieldTag = tags
	return s
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
{{ .Name}} {{ .Reference}}float64 {{ .FieldTag }}
{{- end -}}
`
