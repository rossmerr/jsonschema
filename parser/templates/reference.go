package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Field = (*Reference)(nil)

type Reference struct {
	Types     []string
	Type      *parser.Type
	name      string
	comment   string
	fieldTag  string
	Reference string
}

func NewReference(name, comment string, t *parser.Type, typenames ...string) *Reference {
	return &Reference{
		name:    name,
		comment: comment,
		Type:    t,
		Types:   typenames,
	}
}

func (s *Reference) WithReference(ref bool) parser.Field {
	if s.Type.Kind == parser.Reference || s.Type.Kind == parser.Array {
		return s
	}
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Reference) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Reference) FieldTag() string {
	return s.fieldTag
}

func (s *Reference) Comment() string {
	return s.comment
}

func (s *Reference) Name() string {
	return s.name
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{- $length := len .Types -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ typename .Name}} {{ .Reference}} {{ if eq .Type.Kind.String "array" }}[]{{end}}{{typename .Type.Name}}  {{ .FieldTag }}
{{- end -}}
`
