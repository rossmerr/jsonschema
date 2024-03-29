package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Field = (*String)(nil)

type String struct {
	comment   string
	name      string
	fieldTag  string
	Methods   []string
	Reference string
}

func NewString(name, comment string) *String {
	return &String{
		name:    name,
		comment: comment,
	}
}

func (s *String) WithReference(ref bool) parser.Field {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *String) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *String) FieldTag() string {
	return s.fieldTag
}

func (s *String) Comment() string {
	return s.comment
}

// func (s *String) AppendMethods(methods []string) {
// 	s.Methods = append(s.Methods, methods...)
// }

func (s *String) Name() string {
	return s.name
}

const StringTemplate = `
{{- define "string" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ typename .Name}} {{ .Reference}}string {{ .FieldTag }}


{{- end -}}
`

// {{- range $key, $method := .Methods -}}
// func (s *{{ $.Name }}) {{$method}}(){}
// {{end -}}
