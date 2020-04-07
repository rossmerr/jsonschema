package parser

import "github.com/RossMerr/jsonschema"

type Method struct {
	Comment  string
	Receiver string
	*MethodSignature
	Body string
}

func NewMethod(receiver, name string) *Method {
	return NewMethodFromSignature(receiver, NewMethodSignature(name))
}

func NewMethodFromSignature(receiver string, methodSignature *MethodSignature) *Method {
	return &Method{
		Comment:         "",
		Receiver:        jsonschema.ToTypename(receiver),
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
{{ if .Receiver }}func (s *{{ .Receiver }}){{end}} {{template "methodsignature" .MethodSignature }} { 
{{- .Body -}} 
}
{{- end -}}
`
