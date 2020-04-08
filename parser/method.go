package parser

type Method struct {
	Comment  string
	Receiver string
	*MethodSignature
	Body string
}

func NewMethodFromSignature(receiver string, methodSignature *MethodSignature) *Method {
	return &Method{
		Comment:         "",
		Receiver:        receiver,
		MethodSignature: methodSignature,
		Body:            "",
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

const MethodTemplate = `
{{- define "method" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
{{ if .Receiver }}func (s *{{ typename .Receiver }}){{end}} {{template "methodsignature" .MethodSignature }} { 
{{- .Body -}} 
}
{{- end -}}
`
