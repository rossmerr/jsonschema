package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Receiver = (*Type)(nil)

type Type struct {
	comment  string
	Type     parser.Component
	fieldTag string
	Methods  []*parser.Method
}

func NewType(comment string, t parser.Component) *Type {
	s := &Type{
		comment: comment,
		Type:    t,
	}

	if str, ok := t.(*Struct); ok {
		s.WithMethods(str.UnmarshalStructJSON())
	}

	return s
}

func (s *Type) Comment() string {
	return s.comment
}

func (s *Type) Name() string {
	return emptyString
}

func (s *Type) WithReference(ref bool) parser.Field {
	return s
}

func (s *Type) WithFieldTag(fieldTag string) parser.Field {
	s.fieldTag = fieldTag
	return s
}

func (s *Type) FieldTag() string {
	return s.fieldTag
}

func (s *Type) WithMethods(methods ...*parser.Method) {
	s.Methods = append(s.Methods, methods...)
}

const TypeTemplate = `
{{- define "type" -}}

{{ if .Type.Comment -}}
// {{ .Type.Comment}}
{{else -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{end -}}
type {{template "kind" .Type }} {{ .FieldTag }}
{{range $key, $method := .Methods -}}
	{{ if $method }}
		{{template "method" $method }}
	{{ end}}
{{end }}
{{- end -}}`
