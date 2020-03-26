package parser

import (
	"github.com/RossMerr/jsonschema"
)

type EmbeddedStruct struct {
	name     string
	StructTag  string
	Types []string

}

func NewEmbeddedStruct(typename string, fieldTag string, types ...string) *EmbeddedStruct {
	return &EmbeddedStruct{
		name:     typename,
		StructTag:fieldTag,
		Types: types,
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
{{  .Name }} struct {
	{{range $key, $type := .Types -}}
		*{{ $type}}
	{{end -}}
} {{ .StructTag }}
{{- end -}}
`
