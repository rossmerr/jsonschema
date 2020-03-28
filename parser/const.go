package parser

type Const struct {
	comment   string
	name      string
	Type string
	Value string
}

func NewConst(name, typename, value string ) *Const {
	return &Const{
		name:name,
		Type:typename,
		Value: value,
	}
}

func (s *Const) Comment() string {
	return s.comment
}

func (s *Const) Name() string {
	return s.name
}


const ConstTemplate = `
{{- define "const" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
const {{ title .Name }} {{ .Type }} = "{{ .Value }}"
{{end -}}
`