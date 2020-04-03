package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Reference)(nil)

type Reference struct {
	Type      string
	name      string
	comment   string
	FieldTag  string
	Reference string
}

func NewReference(name, comment, typename string) *Reference {
	return &Reference{
		name:    jsonschema.ToTypename(name),
		comment: comment,
		Type:    typename,
	}
}

func (s *Reference) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Reference) WithReference(ref bool) parser.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Reference) WithFieldTag(tags string) parser.Types {
	s.FieldTag = tags
	return s
}
func (s *Reference) Comment() string {
	return s.comment
}

func (s *Reference) Name() string {
	return s.name
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} {{ .Reference}}{{ .Type}} {{ .FieldTag }}
{{- end -}}
`
