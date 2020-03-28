package parser

import "github.com/RossMerr/jsonschema"

type List []Types

func (s *List) Comment() string {
	return jsonschema.EmptyString
}

func (s *List) Name() string {
	return jsonschema.EmptyString
}


const ListTemplate = `
{{- define "list" -}}
{{range $key, $value := .Values -}}
	{{template "type" $definition  }}
{{end}}
{{end -}}
`