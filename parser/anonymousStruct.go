package parser

import (
	"reflect"

	"github.com/RossMerr/jsonschema"
)

type AnonymousStruct struct {
	comment    string
	Name       string
	id         jsonschema.ID
	Fields     []Types
	StructTag  string
	Interfaces []*Interface
	Method     string
}

func NewAnonymousStruct(ctx *SchemaContext, schema, parent *jsonschema.Schema) *AnonymousStruct {
	fields := []Types{}
	interfaces := []*Interface{}

	for key, propertie := range schema.Properties {
		t := SchemaToType(ctx, key, propertie, schema)
		if propertie.Type() == reflect.Interface {
			i := t.(*Interface)
			interfaces = append(interfaces, i)
		}

		fields = append(fields, t)
	}

	name := schema.ID.Typename()

	structTag := ""
	if parent != nil {
		structTag = ctx.Tags.ToFieldTag(name, schema, parent)
	}

	return &AnonymousStruct{
		comment:    schema.Description,
		Name:       name,
		id:         schema.ID,
		Fields:     fields,
		StructTag:  structTag,
		Interfaces: interfaces,
	}
}

func (s *AnonymousStruct) Comment() string {
	return s.comment
}

func (s *AnonymousStruct) ID() jsonschema.ID {
	return s.id
}

const AnonymousStructTemplate = `
{{- define "struct" -}}
{{  .Name }} struct {
{{range $key, $propertie := .Fields -}}
	{{- if isInterface $propertie -}}
		{{ $propertie.Name }} {{ $propertie.Name }} {{ $propertie.StructTag }}
	{{end -}}
	{{- if isStruct $propertie -}}
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
}

{{ if .Method -}}
	func (s *{{  .Name }}) {{.Method}}(){}
{{end -}}

{{range $key, $interface := .Interfaces -}}
	{{template "interface" $interface }}
{{end}}
{{end}}`
