package parser

import (
	"github.com/RossMerr/jsonschema"
)

type EmbeddedStruct struct {
	Name     string
}

func NewEmbeddedStruct(typename string) *EmbeddedStruct {
	return &EmbeddedStruct{
		Name:     typename,
	}
}

func (s *EmbeddedStruct) Comment() string {
	return jsonschema.EmptyString
}

func (s *EmbeddedStruct) ID() string {
	return jsonschema.EmptyString
}

const EmbeddedStructTemplate = `
{{- define "embeddedStruct" -}}
*{{ .Name}}
{{- end -}}
`
