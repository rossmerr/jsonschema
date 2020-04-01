package types

import "github.com/RossMerr/jsonschema/parser/document"

var _ document.Types = (*Interface)(nil)

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

func (s *Interface) WithReference(ref bool) document.Types {
	return s
}

func (s *Interface) WithFieldTag(tags string) document.Types {
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
