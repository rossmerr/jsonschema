package types

import (
	"fmt"
	"strings"

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

func HandleEnum(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	reference := "*"
	if jsonschema.Contains(schema.Required, strings.ToLower(name)) {
		reference = ""
	}

	constItems := []*ConstItem{}
	for _, value := range schema.Enum {
		c := ConstItem{
			Name:  jsonschema.ToTypename(value),
			Type:  name,
			Value: value,
		}
		constItems = append(constItems, &c)
	}
	c := NewConst(constItems...)
	typenameEnum := name + "Enum"
	if _, ok := ctx.Globals[typenameEnum]; !ok {
		ctx.Globals[typenameEnum] = c
	} else {
		return nil, fmt.Errorf("enum, global keys need to be unique found %v more than once, in %v", name, schema.Parent.ID)
	}

	return &Enum{
		comment:   schema.Description,
		name:      name,
		Type:      schema.Type.String(),
		Reference: reference,
		Values:    schema.Enum,
		items:     constItems,
	}, nil
}

func GlobalEnum(ctx *document.DocumentContext, enum *Enum, name string) document.Types {
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
