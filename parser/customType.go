package parser

type CustomType struct {
	comment string
	name    string
	Type    string
}

func NewCustomType(name, typename string) *CustomType {
	return &CustomType{
		name: name,
		Type: typename,
	}
}

func (s *CustomType) Comment() string {
	return s.comment
}

func (s *CustomType) Name() string {
	return s.name
}

const CustomTypeTemplate = `
{{- define "customtype" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
 type {{ .Name }} {{ .Type }}
{{end}}
`
