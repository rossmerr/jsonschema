package parser

type Type struct {
	comment string
	Type    Types
	Methods []string
}

func PrefixType(t Types, methods ...string) *Type {
	return &Type{
		comment: t.Comment(),
		Methods: methods,
		Type:    t,
	}
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
