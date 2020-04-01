package types

import "github.com/RossMerr/jsonschema/parser/document"

var _ document.Types = (*Type)(nil)

// Obsolete
type Type struct {
	comment string
	Type    document.Types
	Methods []string
}

// Obsolete
func PrefixType(t document.Types, methods ...string) *Type {
	return &Type{
		comment: t.Comment(),
		Methods: methods,
		Type:    t,
	}
}

func (s *Type) WithReference(ref bool) document.Types {
	return s
}

func (s *Type) WithFieldTag(tags string) document.Types {
	return s
}

func (s *Type) Comment() string {
	return s.comment
}

func (s *Type) Name() string {
	return s.Type.Name()
}

const TypeTemplate = `
{{- define "type" -}}
{{ if .Type.Comment -}}
// {{.Type.Comment}}
{{end -}}
type {{template "kind" .Type }}

{{range $key, $method := .Methods -}}
func (s {{ $.Name }}) {{$method}}(){}
{{end }}
{{- end -}}
`
