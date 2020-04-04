package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*List)(nil)

type List struct {
	Items []parser.Types
}

func (s *List) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *List) WithReference(ref bool) parser.Types {
	return s
}

func (s *List) WithFieldTag(tags string) parser.Types {
	return s
}

func (s *List) FieldTag() string {
	return jsonschema.EmptyString
}

func (s *List) Comment() string {
	return ""
}

func (s *List) Name() string {
	return ""
}
