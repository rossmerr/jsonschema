package parser

type Integer struct {
	comment    string
	name       string
	Validation string
	FieldTag   string
	Reference  string
}

func NewInteger(name *Name, description, fieldTag string, isReference bool) *Integer {
	reference := ""
	if isReference {
		reference = "*"
	}

	return &Integer{
		comment:   description,
		name:      name.Fieldname(),
		FieldTag:  fieldTag,
		Reference: reference,
	}
}

func (s *Integer) Comment() string {
	return s.comment
}

func (s *Integer) Name() string {
	return s.name
}

const IntegerTemplate = `
{{- define "integer" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} {{ .Reference}}int32 {{ .FieldTag }}
{{- end -}}
`
