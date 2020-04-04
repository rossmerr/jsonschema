package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*AnyOf)(nil)

type AnyOf struct {
	*Reference
}

func (s *AnyOf) WithMethods(methods ...*parser.Method) parser.Types {
	return s.Reference.WithMethods(methods...)
}

func (s *AnyOf) WithReference(ref bool) parser.Types {
	return s.Reference.WithReference(ref)
}

func (s *AnyOf) WithFieldTag(tags string) parser.Types {
	return s.Reference.WithFieldTag(tags)
}

func (s *AnyOf) FieldTag() string {
	return s.Reference.FieldTag()
}

func (s *AnyOf) Comment() string {
	return s.Reference.Comment()
}

func (s *AnyOf) Name() string {
	return s.Reference.Name()
}

const AnyOfTemplate = `
{{- define "anyof" -}}
{{template "kind" .Reference }}
{{- end -}}`
