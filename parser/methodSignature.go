package parser

type MethodSignature struct {
	Name    string
	Inputs  []*Parameter
	Outputs []*Parameter
}

func NewMethodSignature(name string) *MethodSignature {
	return &MethodSignature{
		Name:    name,
		Inputs:  []*Parameter{},
		Outputs: []*Parameter{},
	}
}

func (s *MethodSignature) WithInputs(inputs ...*Parameter) *MethodSignature {
	s.Inputs = append(s.Inputs, inputs...)
	return s
}

func (s *MethodSignature) WithOutputs(outputs ...*Parameter) *MethodSignature {
	s.Outputs = append(s.Outputs, outputs...)
	return s
}

const MethodSignatureTemplate = `
{{- define "methodsignature" -}}
{{ .Name -}}(
{{- range $key, $input := .Inputs -}}
{{ $input.Name }} {{ $input.Type }}
{{- end -}}
) {{- if .Outputs -}}
(
{{- range $key, $output := .Outputs -}}
{{ $output.Name }} {{ $output.Type }}
{{- end -}}
) {{ end }}
{{- end -}}
`
