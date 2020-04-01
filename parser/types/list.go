package types

import "github.com/RossMerr/jsonschema/parser/document"

var _ document.Types = (*List)(nil)

type List struct {
	Items []document.Types
}

func (s *List) WithReference(ref bool) document.Types {
	return s
}

func (s *List) WithFieldTag(tags string) document.Types {
	return s
}

func (s *List) Comment() string {
	return ""
}

func (s *List) Name() string {
	return ""
}
