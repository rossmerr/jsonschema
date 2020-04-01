package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Boolean)(nil)

type Boolean struct {
	comment   string
	name      string
	FieldTag  string
	Reference string
}

func HandleBoolean(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	return &Boolean{
		comment: schema.Description,
		name:    jsonschema.ToTypename(name),
	}, nil
}

func (s *Boolean) WithReference(ref bool) document.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Boolean) WithFieldTag(tags string) document.Types {
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
