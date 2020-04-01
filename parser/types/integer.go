package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Integer)(nil)

type Integer struct {
	comment    string
	name       string
	Validation string
	FieldTag   string
	Reference  string
}

func HandleInteger(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	return &Integer{
		comment: schema.Description,
		name:    jsonschema.ToTypename(name),
	}, nil
}

func (s *Integer) WithReference(ref bool) document.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Integer) WithFieldTag(tags string) document.Types {
	s.FieldTag = tags
	return s
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
