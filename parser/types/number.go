package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Number)(nil)

type Number struct {
	comment    string
	name       string
	Validation string
	FieldTag   string
	Reference  string
}

func HandleNumber(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	return &Number{
		comment: schema.Description,
		name:    jsonschema.ToTypename(name),
	}, nil
}

func (s *Number) WithReference(ref bool) document.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Number) WithFieldTag(tags string) document.Types {
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
