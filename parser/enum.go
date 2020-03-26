package parser

import (

)

type Enum struct {
	comment  string
	name     string
	Values   []string
	FieldTag string
	Reference  string
}

func NewEnum(name *Name, description, fieldTag string, isReference bool, values []string) *Enum {
	reference := ""
	if isReference {
		reference = "*"
	}

	return &Enum{
		comment:  description,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
		Reference:  reference,
		Values:  values,
	}
}

func (s *Enum) Comment() string {
	return s.comment
}

func (s *Enum) Name() string {
	return s.name
}

const EnumTemplate = `
{{- define "enum" -}}


// test
{{end -}}
`
