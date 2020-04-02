package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*InterfaceReference)(nil)

type InterfaceReference struct {
	Type     string
	name     string
	FieldTag string
}

func NewInterfaceReference(name, typename string) *InterfaceReference {
	return &InterfaceReference{
		name: jsonschema.ToTypename(name),
		Type: typename,
	}
}
func (s *InterfaceReference)	WithMethods(methods ...string) parser.Types {
	return s
}

func (s *InterfaceReference) WithReference(ref bool) parser.Types {
	return s
}

func (s *InterfaceReference) WithFieldTag(tags string) parser.Types {
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
