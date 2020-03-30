package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Struct struct {
	comment   string
	name      string
	Fields    []Types
	StructTag string
}

func NewStruct(ctx *SchemaContext, name *Name, properties map[string]*jsonschema.Schema, comment, fieldTag string, required ...string) (Types, error) {

	fields := []Types{}
	for key, propertie := range properties {
		s, err := schemaToType(ctx, NewName(key), propertie, true, required...)
		if err != nil {
			return nil, err
		}
		fields = append(fields, s)
	}

	return &Struct{
		comment:   comment,
		name:      name.Fieldname(),
		Fields:    fields,
		StructTag: fieldTag,
	}, nil
}

func (s *Struct) Comment() string {
	return s.comment
}

func (s *Struct) Name() string {
	return s.name
}

const StructTemplate = `
{{- define "struct" -}}
{{  .Name }} struct {
{{range $key, $propertie := .Fields -}}
	{{template "kind" $propertie }}
{{end -}}
} {{ .StructTag }}

{{end}}`
