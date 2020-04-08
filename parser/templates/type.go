package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Type)(nil)

type Type struct {
	comment  string
	Type     parser.Types
	fieldTag string
	Methods  []*parser.Method
}

func (s *Type) Comment() string {
	return s.comment
}

func (s *Type) Name() string {
	return EmptyString
}

func (s *Type) WithFieldTag(fieldTag string) parser.Types {
	s.fieldTag = fieldTag
	return s
}

func (s *Type) FieldTag() string {
	return s.fieldTag
}

func (s *Type) WithReference(bool) parser.Types {
	return s
}

func (s *Type) WithMethods(methods ...*parser.Method) parser.Types {
	s.Methods = append(s.Methods, methods...)
	return s
}

func NewType(comment string, t parser.Types) *Type {
	return &Type{
		comment: comment,
		Type:    t,
	}
}

const TypeTemplate = `
{{- define "type" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
type {{template "kind" .Type }} {{ .FieldTag }}
{{range $key, $method := .Methods -}}
	{{ if $method }}
		{{template "method" $method }}
	{{ end}}
{{end }}
{{- end -}}`
