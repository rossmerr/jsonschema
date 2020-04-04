package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*OneOf)(nil)

type OneOf struct {
	*Reference
}

func (s *OneOf) WithMethods(methods ...*parser.Method) parser.Types {
	return s.Reference.WithMethods(methods...)
}

func (s *OneOf) WithReference(ref bool) parser.Types {
	return s.Reference.WithReference(ref)
}

func (s *OneOf) WithFieldTag(tags string) parser.Types {
	return s.Reference.WithFieldTag(tags)
}

func (s *OneOf) FieldTag() string {
	return s.Reference.FieldTag()
}

func (s *OneOf) Comment() string {
	return s.Reference.Comment()
}

func (s *OneOf) Name() string {
	return s.Reference.Name()
}

const OneOfTemplate = `
{{- define "oneof" -}}
{{template "kind" .Reference }}
{{- end -}}`
