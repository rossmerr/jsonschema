package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*EmbeddedStruct)(nil)

type EmbeddedStruct struct {
	name      string
	StructTag string
	Types     []string
}

func NewEmbeddedStruct(name string, fieldTag string, types ...string) *EmbeddedStruct {
	return &EmbeddedStruct{
		name:      name,
		StructTag: fieldTag,
		Types:     types,
	}
}

func (s *EmbeddedStruct) WithReference(ref bool) document.Types {
	return s
}

func (s *EmbeddedStruct) WithFieldTag(tags string) document.Types {
	return s
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
