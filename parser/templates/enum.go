package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Field = (*Enum)(nil)

type Enum struct {
	comment   string
	name      string
	Type      string
	Values    []string
	fieldTag  string
	Reference string
	items     []*ConstItem
}

func NewEnum(name, comment, typename string, values []string, items []*ConstItem) *Enum {
	return &Enum{
		comment: comment,
		name:    name,
		Type:    typename,
		Values:  values,
		items:   items,
	}
}

func (s *Enum) WithReference(ref bool) parser.Field {
	return s
}

func (s *Enum) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Enum) FieldTag() string {
	return s.fieldTag
}

func (s *Enum) Comment() string {
	return s.comment
}

func (s *Enum) Name() string {
	return s.name
}

const EnumTemplate = `
{{- define "enum" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
type {{ typename .Name}} {{ .Type }} 
{{end -}}
`
