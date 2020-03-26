package parser

type Array struct {
	comment  string
	name     string
	Type     string
	FieldTag string
}

func NewArray(name *Name, description, fieldTag, arrayType string) *Array {
	return &Array{
		comment:  description,
		name:     name.Fieldname(),
		Type:     arrayType,
		FieldTag: fieldTag,
	}
}

func (s *Array) Comment() string {
	return s.comment
}

func (s *Array) Name() string {
	return s.name
}

const ArrayTemplate = `
{{- define "array" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} []{{ .Type }} {{ .FieldTag }}
{{- end -}}
`
