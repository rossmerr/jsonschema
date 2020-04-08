package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*List)(nil)

type List struct {
	Items []parser.Component
}

func (s *List) Name() string {
	return ""
}
