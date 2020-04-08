package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*Type)(nil)
var _ parser.Field = (*Type)(nil)
var _ parser.Receiver = (*Type)(nil)

type Type struct {
	comment  string
	Type     parser.Component
	fieldTag string
	Methods  []*parser.Method
}

func (s *Type) Comment() string {
	return s.comment
}

func (s *Type) Name() string {
	return EmptyString
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

func NewType(comment string, t parser.Component) *Type {
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
