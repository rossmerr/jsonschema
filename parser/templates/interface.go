package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*Interface)(nil)

type Interface struct {
	comment                string
	CommentImplementations string
	name                   string
	MethodSignatures       []*parser.MethodSignature
}

func NewInterface(typename string) *Interface {
	return &Interface{
		name: typename,
	}
}

func (s *Interface) WithMethodSignature(methodSignature ...*parser.MethodSignature) *Interface {
	s.MethodSignatures = append(s.MethodSignatures, methodSignature...)
	return s
}

func (s *Interface) Comment() string {
	return s.comment
}

func (s *Interface) Name() string {
	return s.name
}

const InterfaceTemplate = `
{{- define "interface" -}}
{{if .Comment -}}
// {{ .Comment}}
{{ else -}}
// {{ typename .Name }}
{{end -}}
{{if .CommentImplementations -}}
// {{ .CommentImplementations}}
{{end -}}
type {{ typename .Name }} interface {	
	{{range $key, $method := .MethodSignatures -}}
		{{template "methodsignature" $method }}
	{{end }}
}
{{end -}}
`
