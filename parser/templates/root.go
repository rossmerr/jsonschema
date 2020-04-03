package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Root)(nil)

type Root struct {
	comment string
	Type    parser.Types
	Methods []*parser.Method
}

func NewRoot(comment string, t parser.Types) parser.Types {
	return &Root{
		comment: comment,
		Type:    t,
		Methods: []*parser.Method{},
	}
}

func (s *Root) WithMethods(methods ...*parser.Method) parser.Types {
	s.Methods = append(s.Methods, methods...)
	return s
}
func (s *Root) WithReference(ref bool) parser.Types {
	return s
}

func (s *Root) WithFieldTag(tags string) parser.Types {
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
	{{template "method" $method }}
{{end }}
{{- end -}}
`
