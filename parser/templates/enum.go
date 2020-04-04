package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Enum)(nil)

type Enum struct {
	comment   string
	name      string
	Type      string
	Values    []string
	fieldTag  string
	Reference string
	items     []*ConstItem
}

func NewEnum(name, comment, typename string, values []string, items []*ConstItem) parser.Types {
	return &Enum{
		comment: comment,
		name:    name,
		Type:    typename,
		Values:  values,
		items:   items,
	}
}
func (s *Enum) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Enum) WithReference(ref bool) parser.Types {
	return s
}

func (s *Enum) WithFieldTag(tags string) parser.Types {
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
{{ .Name}} {{ .Type }} 
{{end -}}
`
