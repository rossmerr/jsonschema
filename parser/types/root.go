package types

import (
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Root)(nil)

type Root struct {
	comment string
	Type    document.Types
	Methods []string
}

func NewRoot(comment string, t document.Types) document.Types {
	return &Root{
		comment: comment,
		Type:    t,
	}
}

func (s *Root) WithMethods(methods ...string) document.Types {
	s.Methods = methods
	return s
}
func (s *Root) WithReference(ref bool) document.Types {
	return s
}

func (s *Root) WithFieldTag(tags string) document.Types {
	return s
}

func (s *Root) Comment() string {
	return s.comment
}

func (s *Root) Name() string {
	return s.Type.Name()
}

const RootTemplate = `
{{- define "root" -}}
{{ if .Type.Comment -}}
// {{.Type.Comment}}
{{end -}}
type {{template "kind" .Type }}

{{range $key, $method := .Methods -}}
func (s {{ $.Name }}) {{$method}}(){}
{{end }}
{{- end -}}
`
