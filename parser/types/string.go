package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*String)(nil)

type String struct {
	comment   string
	name      string
	FieldTag  string
	Methods   []string
	Reference string
}

func HandleString(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	return &String{
		comment: schema.Description,
		name:    jsonschema.ToTypename(name),
	}, nil
}

func (s *String) WithReference(ref bool) document.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *String) WithFieldTag(tags string) document.Types {
	s.FieldTag = tags
	return s
}

func (s *String) Comment() string {
	return s.comment
}

func (s *String) AppendMethods(methods []string) {
	s.Methods = append(s.Methods, methods...)
}

func (s *String) Name() string {
	return s.name
}

const StringTemplate = `
{{- define "string" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Reference}}string {{ .FieldTag }}

{{- range $key, $method := .Methods -}}
	func (s *{{ $.Name }}) {{$method}}(){}
{{end -}}
{{- end -}}
`
