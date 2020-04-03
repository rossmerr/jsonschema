package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*AllOf)(nil)

type AllOf struct {
	*Struct
}

func (s *AllOf) WithMethods(methods ...*parser.Method) parser.Types {
	return s.Struct.WithMethods(methods...)
}

func (s *AllOf) WithReference(ref bool) parser.Types {
	return s.Struct.WithReference(ref)
}

func (s *AllOf) WithFieldTag(tags string) parser.Types {
	return s.Struct.WithFieldTag(tags)
}

func (s *AllOf) Comment() string {
	return s.Struct.Comment()
}

func (s *AllOf) Name() string {
	return s.Struct.Name()
}

const AllOfTemplate = `
{{- define "allof" -}}
{{template "kind" .Struct }}
{{- end -}}`
