package parser

type Boolean struct {
	comment  string
	name     string
	FieldTag string
	Reference string

}

func NewBoolean(name *Name, description, fieldTag string, isReference bool) *Boolean {
	reference := ""
	if isReference {
		reference = "*"
	}

	return &Boolean{
		comment:  description,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
		Reference:  reference,
	}
}

func (s *Boolean) Comment() string {
	return s.comment
}

func (s *Boolean) Name() string {
	return s.name
}


const BooleanTemplate = `
{{- define "boolean" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Reference}}bool {{ .FieldTag }}
{{- end -}}
`
