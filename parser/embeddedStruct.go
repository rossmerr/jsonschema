package parser

import (
	"github.com/RossMerr/jsonschema"
)

type EmbeddedStruct struct {
	name     string
}

func NewEmbeddedStruct(typename string) *EmbeddedStruct {
	return &EmbeddedStruct{
		name:     typename,
	}
}

func (s *EmbeddedStruct) Comment() string {
	return jsonschema.EmptyString
}

func (s *EmbeddedStruct) Name() string {
	return s.name
}


const EmbeddedStructTemplate = `
{{- define "embeddedStruct" -}}
*{{ .Name}}
{{- end -}}
`
