package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*InterfaceReference)(nil)

type InterfaceReference struct {
	Type     string
	name     string
	fieldTag string
}

func NewInterfaceReference(name, typename string) *InterfaceReference {
	return &InterfaceReference{
		name: jsonschema.ToTypename(name),
		Type: typename,
	}
}
func (s *InterfaceReference) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *InterfaceReference) WithReference(ref bool) parser.Types {
	return s
}

func (s *InterfaceReference) WithFieldTag(tags string) parser.Types {
	s.fieldTag = tags
	return s
}

func (s *InterfaceReference) FieldTag() string {
	return s.fieldTag
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
