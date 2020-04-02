package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Enum)(nil)

type Enum struct {
	comment   string
	name      string
	Type      string
	Values    []string
	FieldTag  string
	Reference string
	items     []*ConstItem
}

func NewEnum(name, comment, typename string, values []string, items []*ConstItem) document.Types {
	return &Enum{
		comment: comment,
		name:    jsonschema.ToTypename(name),
		Type:    typename,
		Values:  values,
		items:   items,
	}
}

func GlobalEnum(ctx *document.Document, enum *Enum, name string) document.Types {
	typename := name + "_" + enum.name
	enum.name = typename
	for _, item := range enum.items {
		item.Type = typename
	}
	ctx.Globals[typename] = PrefixType(enum)
	ref := &Reference{
		Type:      enum.name,
		FieldTag:  enum.FieldTag,
		Reference: enum.Reference,
	}
	ref.name = enum.name

	return ref
}

func (s *Enum) WithReference(ref bool) document.Types {
	return s
}

func (s *Enum) WithFieldTag(tags string) document.Types {
	s.FieldTag = tags
	return s
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
