package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Struct struct {
	Package    string
	Comment    string
	Name       string
	ID         jsonschema.ID
	Properties []Types
	StructTag  string
}

func NewStruct(ctx *SchemaContext, key string, schema, parent *jsonschema.Schema) *Struct {
	properties := []Types{}

	for key, propertie := range schema.Properties {
		properties = append(properties, SchemaToType(ctx, key, propertie, schema))
	}

	name := schema.ID.Typename()
	if name == "" {
		name = key
	}

	structTag := ""
	if parent != nil {
		structTag = ctx.Tags.ToFieldTag(strings.Title(name), schema, parent)
	}

	return &Struct{
		Package:    ctx.Package,
		Comment:    schema.Description,
		Name:       strings.Title(name),
		ID:         schema.ID,
		Properties: properties,
		StructTag:  structTag,
	}
}

func (s *Struct) types() {}

const StructTemplate = `
{{- define "struct" -}}
{{  .Name }} struct {
{{range $key, $propertie := .Properties -}}
	{{- if isStruct $propertie -}}string
		{{template "struct" $propertie }}
	{{end -}}
	{{- if isArray $propertie -}}
		{{template "array" $propertie }}
	{{end -}}
	{{- if isNumber $propertie -}}
		{{template "number" $propertie }}
	{{end -}}
	{{- if isString $propertie -}}
		{{template "string" $propertie }}
	{{end -}}
	{{- if isBoolean $propertie -}}
		{{template "boolean" $propertie }}
	{{end -}}
{{end -}}
} {{ .StructTag }}
{{end}}`
