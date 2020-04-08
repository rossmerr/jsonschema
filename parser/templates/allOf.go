package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Component = (*AllOf)(nil)

type AllOf struct {
	*Reference
}

func (s *AllOf) Name() string {
	return s.Reference.Name()
}

const AllOfTemplate = `
{{- define "allof" -}}
{{template "kind" .Reference }}
{{- end -}}`
