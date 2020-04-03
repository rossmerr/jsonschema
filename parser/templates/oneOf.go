package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*OneOf)(nil)

type OneOf struct {
	*InterfaceReference
}

func (s *OneOf) WithMethods(methods ...*parser.Method) parser.Types {
	return s.InterfaceReference.WithMethods(methods...)
}

func (s *OneOf) WithReference(ref bool) parser.Types {
	return s.InterfaceReference.WithReference(ref)
}

func (s *OneOf) WithFieldTag(tags string) parser.Types {
	return s.InterfaceReference.WithFieldTag(tags)
}

func (s *OneOf) FieldTag() string {
	return s.InterfaceReference.FieldTag()
}


func (s *OneOf) Comment() string {
	return s.InterfaceReference.Comment()
}

func (s *OneOf) Name() string {
	return s.InterfaceReference.Name()
}

const OneOfTemplate = `
{{- define "oneof" -}}
{{template "kind" .InterfaceReference }}
{{- end -}}`
