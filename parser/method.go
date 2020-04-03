package parser

type Method struct {
	Comment  string
	Receiver string
	Name     string
	Inputs   []*Parameter
	Outputs  []*Parameter
	Body     string
}

func NewMethod(receiver, name string) *Method {
	return &Method{
		Comment:  "",
		Receiver: receiver,
		Name:     name,
		Inputs:   []*Parameter{},
		Outputs:  []*Parameter{},
		Body:     "",
	}
}

func (s *Method) WithInputs(inputs ...*Parameter) *Method {
	s.Inputs = append(s.Inputs, inputs...)
	return s
}

func (s *Method) WithOutputs(outputs ...*Parameter) *Method {
	s.Outputs = append(s.Outputs, outputs...)
	return s
}

type Parameter struct {
	Comment string
	Name    string
	Type    string
}

func NewParameter(name, typename string) *Parameter {
	return &Parameter{
		Name: name,
		Type: typename,
	}
}

const MethodTemplate = `
{{- define "method" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
func (s *{{ .Receiver }}) {{ .Name -}}
(
{{- range $key, $input := .Inputs -}}
	{{ $input.Name }} {{ $input.Type }}
{{- end -}}
)
{{- if .Outputs -}}
(
{{- range $key, $output := .Outputs -}}
	{{ $output.Name }} {{ $output.Type }}
{{- end -}}
)
{{- end -}}{ 
{{ .Body }} 
}
{{- end -}}
`
