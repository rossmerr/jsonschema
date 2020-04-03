package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*AnyOf)(nil)

type AnyOf struct {
	*InterfaceReference
}

func (s *AnyOf) WithMethods(methods ...*parser.Method) parser.Types {
	return s.InterfaceReference.WithMethods(methods...)
}

func (s *AnyOf) WithReference(ref bool) parser.Types {
	return s.InterfaceReference.WithReference(ref)
}

func (s *AnyOf) WithFieldTag(tags string) parser.Types {
	return s.InterfaceReference.WithFieldTag(tags)
}

func (s *AnyOf) FieldTag() string {
	return s.InterfaceReference.FieldTag()
}


func (s *AnyOf) Comment() string {
	return s.InterfaceReference.Comment()
}

func (s *AnyOf) Name() string {
	return s.InterfaceReference.Name()
}

const AnyOfTemplate = `
{{- define "anyof" -}}
{{template "kind" .InterfaceReference }}
{{- end -}}`
