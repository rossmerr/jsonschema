package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*AnyOf)(nil)

type AnyOf struct {
	*Reference
}

func (s *AnyOf) Name() string {
	return s.Reference.Name()
}

const AnyOfTemplate = `
{{- define "anyof" -}}
{{template "kind" .Reference }}
{{- end -}}`
