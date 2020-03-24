package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Interface struct {
	comment string
	CommentImplementations string
	Name    string

}

func NewInterface(typename string) *Interface {
	return &Interface{
		Name:    typename,
	}
}

func (s *Interface) Comment() string {
	return s.comment
}

func (s *Interface) ID() string {
	return jsonschema.EmptyString
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
