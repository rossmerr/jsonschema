package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*AllOf)(nil)

type AllOf struct {
	*Reference
}

func (s *AllOf) WithMethods(methods ...*parser.Method) parser.Types {
	return s.Reference.WithMethods(methods...)
}

func (s *AllOf) WithReference(ref bool) parser.Types {
	return s.Reference.WithReference(ref)
}

func (s *AllOf) WithFieldTag(tags string) parser.Types {
	return s.Reference.WithFieldTag(tags)
}

func (s *AllOf) Comment() string {
	return s.Reference.Comment()
}

func (s *AllOf) Name() string {
	return s.Reference.Name()
}

const AllOfTemplate = `
{{- define "allof" -}}
{{template "kind" .Reference }}
{{- end -}}`
