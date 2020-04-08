package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*OneOf)(nil)

type OneOf struct {
	*Reference
}

func (s *OneOf) Name() string {
	return s.Reference.Name()
}

const OneOfTemplate = `
{{- define "oneof" -}}
{{template "kind" .Reference }}
{{- end -}}`
