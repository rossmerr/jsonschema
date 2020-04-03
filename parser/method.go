package parser

type Method struct {
	Comment  string
	Receiver string
	Name     string
	Inputs   []string
	Outputs  []string
	Body     string
}

func NewMethod(receiver, name string) *Method {
	return &Method{
		Comment:  "",
		Receiver: receiver,
		Name:     name,
		Inputs:   []string{},
		Outputs:  []string{},
		Body:     "",
	}
}

func (s *Method) WithInputs(inputs ...string) *Method {
	s.Inputs = append(s.Inputs, inputs...)
	return s
}

func (s *Method) WithOutputs(outputs ...string) *Method {
	s.Outputs = append(s.Outputs, outputs...)
	return s
}

const MethodTemplate = `
{{- define "method" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
func (s *{{ .Receiver }}) {{ .Name -}}
(
{{- range $key, $input := .Inputs -}}
	{{ $input }}
{{end -}}
)
{{- range $key, $output := .Outputs -}}
	{{ $output }}
{{end -}}
{ {{- .Body -}} }
{{- end -}}
`
