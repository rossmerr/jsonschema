package parser

type CustomType struct {
	comment string
	Type    Types
	Methods []string
}

func PrefixType(t Types, methods ...string) *CustomType {
	return &CustomType{
		comment: t.Comment(),
		Methods: methods,
		Type:    t,
	}
}

func (s *CustomType) Comment() string {
	return s.comment
}

func (s *CustomType) Name() string {
	return s.Type.Name()
}

const CustomTypeTemplate = `
{{- define "customtype" -}}
{{ if .Type.Comment -}}
// {{.Type.Comment}}
{{end -}}
type {{template "kind" .Type }}

{{range $key, $method := .Methods -}}
func (s {{ $.Name }}) {{$method}}(){}
{{end }}
{{- end -}}
`
