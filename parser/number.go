package parser

type Number struct {
	comment    string
	name       string
	Validation string
	FieldTag   string
	Reference string
}

func NewNumber(name *Name, description, fieldTag string, isReference bool) *Number {
	reference := ""
	if isReference {
		reference = "*"
	}

	return &Number{
		comment:   description,
		name:      name.Fieldname(),
		FieldTag:  fieldTag,
		Reference: reference,
	}
}

func (s *Number) Comment() string {
	return s.comment
}

func (s *Number) Name() string {
	return s.name
}


const NumberTemplate = `
{{- define "number" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Reference}}float64 {{ .FieldTag }}
{{- end -}}
`
