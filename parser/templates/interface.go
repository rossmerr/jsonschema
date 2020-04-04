package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Interface)(nil)

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

func (s *Interface) WithMethodSignature(methodSignature ...*parser.MethodSignature) parser.Types {
	s.MethodSignatures = append(s.MethodSignatures, methodSignature...)
	return s
}

func (s *Interface) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Interface) WithReference(ref bool) parser.Types {
	return s
}

func (s *Interface) WithFieldTag(tags string) parser.Types {
	return s
}

func (s *Interface) FieldTag() string {
	return jsonschema.EmptyString
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
// {{ .Name }}
{{end -}}
{{if .CommentImplementations -}}
// {{ .CommentImplementations}}
{{end -}}
type {{ .Name }} interface {	
	{{range $key, $method := .MethodSignatures -}}
		{{template "methodsignature" $method }}
	{{end }}
}
{{end -}}
`
