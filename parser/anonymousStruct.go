package parser

import (
	"reflect"

	"github.com/RossMerr/jsonschema"
)

type AnonymousStruct struct {
	comment    string
	Name       string
	id         string
	Fields     []Types
	StructTag  string
	Interfaces []*Interface
	Method     string
}

func NewAnonymousStruct(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *AnonymousStruct {
	fields := []Types{}
	interfaces := []*Interface{}

	for key, propertie := range schema.Properties {
		t := SchemaToType(WrapContext(ctx, schema), key, propertie, schema.Required)
		if propertie.Type() == reflect.Interface {
			if i, ok := t.(*Interface); ok {
				interfaces = append(interfaces, i)
			}
		}

		fields = append(fields, t)
	}

	return &AnonymousStruct{
		comment:    schema.Description,
		Name:       typename,
		id:         schema.ID.String(),
		Fields:     fields,
		StructTag:  ctx.Tags.ToFieldTag(typename, schema, required),
		Interfaces: interfaces,
	}
}

func (s *AnonymousStruct) Comment() string {
	return s.comment
}

func (s *AnonymousStruct) ID() string {
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
