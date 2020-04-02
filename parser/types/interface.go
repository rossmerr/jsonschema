package types

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Interface)(nil)

type Interface struct {
	comment                string
	CommentImplementations string
	name                   string
}

func NewInterface(typename string) *Interface {
	return &Interface{
		name: typename,
	}
}

func (s *Interface)	WithMethods(methods ...string) parser.Types {
	return s
}

func (s *Interface) WithReference(ref bool) parser.Types {
	return s
}

func (s *Interface) WithFieldTag(tags string) parser.Types {
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
// {{ .Name }}
{{end -}}
{{if .CommentImplementations -}}
// {{ .CommentImplementations}}
{{end -}}
type {{ .Name }} interface {
	{{ .Name}}()
}
{{end -}}
`
