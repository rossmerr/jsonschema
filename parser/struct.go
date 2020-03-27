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

func NewStruct(ctx *SchemaContext, name *Name, properties map[string]*jsonschema.Schema, comment, fieldTag string, required ...string) Types {

	fields := []Types{}
	for key, propertie := range properties {
		fields = append(fields, schemaToType(ctx, NewName(key), propertie, true, required...))
	}

	return &Struct{
		comment:   comment,
		name:      name.Fieldname(),
		Fields:    fields,
		StructTag: fieldTag,
	}
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
