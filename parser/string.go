package parser

type String struct {
	comment  string
	name     string
	FieldTag string
	Methods    []string
}

func NewString(name *Name, description, fieldTag string) *String {
	return &String{
		comment:  description,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}
}

func (s *String) Comment() string {
	return s.comment
}

func (s *String) AppendMethods(methods []string) {
	s.Methods = append(s.Methods, methods...)
}

func (s *String) Name() string {
	return s.name
}

const StringTemplate = `
{{- define "string" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ .Name}} string {{ .FieldTag }}

{{range $key, $method := .Methods -}}
	func (s *{{ $.Name }}) {{$method}}(){}
{{end -}}
{{- end -}}
`
