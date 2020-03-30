package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
)

type Enum struct {
	comment   string
	name      string
	Type      string
	Values    []string
	FieldTag  string
	Reference string
}

func NewEnum(ctx *SchemaContext, name *Name, description, fieldTag string, isReference bool, values []string) (*Enum, error) {
	reference := ""
	if isReference {
		reference = "*"
	}
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname() + "Enum"

	list := List{
		NewCustomType(typename, "string"),
	}

	for _, value := range values {
		c := NewConst(jsonschema.ToTypename(value), typename, value)
		list = append(list, c)
	}

	if _, ok := ctx.Globals[typename]; !ok {
		ctx.Globals[typename] = &list
	} else {
		return nil, fmt.Errorf("Global keys need to be unique found %v more than once, in %v", typename, parent.ID)
	}

	return &Enum{
		comment:   description,
		name:      name.Fieldname(),
		Type:      typename,
		FieldTag:  fieldTag,
		Reference: reference,
		Values:    values,
	}, nil
}

func (s *Enum) Comment() string {
	return s.comment
}

func (s *Enum) Name() string {
	return s.name
}

const EnumTemplate = `
{{- define "enum" -}}
{{ .Name}} {{ .Type }} {{ .FieldTag }}
{{end -}}
`

// {{ if .Comment -}}
// // {{.Comment}}
// {{end -}}
