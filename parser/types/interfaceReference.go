package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*InterfaceReference)(nil)

type InterfaceReference struct {
	Type     string
	name     string
	FieldTag string
}

func (s *InterfaceReference) WithReference(ref bool) document.Types {
	return s
}

func (s *InterfaceReference) WithFieldTag(tags string) document.Types {
	s.FieldTag = tags
	return s
}

func (s *InterfaceReference) Comment() string {
	return jsonschema.EmptyString
}

func (s *InterfaceReference) Name() string {
	return s.name
}

const InterfaceReferenceTemplate = `
{{- define "interfacereference" -}}
{{ .Name}} {{ .Type}} {{ .FieldTag }}
{{- end -}}
`
