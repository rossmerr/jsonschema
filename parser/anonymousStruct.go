package parser

import (
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
	typename = jsonschema.Structname(typename)

	fields = addProperties(ctx, schema, fields)
	for key, child := range schema.AllOf {
		if key == "$ref" {
			_, typename, _ := ResolvePointer(ctx, child.Ref)
			fields = append(fields, NewEmbeddedStruct(typename))
			continue
		}

		if child.Ref != jsonschema.EmptyString {
			_, typename, _ := ResolvePointer(ctx, child.Ref)
			fields = append(fields, NewReference(typename, jsonschema.Fieldname(key)))
			continue
		}

		fields = addProperties(ctx, child, fields)
	}

	for _, t := range fields {
		if ref, ok := t.(*InterfaceReference); ok {
			i := NewInterface(ref.Type)
			interfaces = append(interfaces, i)
		}
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

func addProperties(ctx *SchemaContext, schema *jsonschema.Schema, fields []Types) ([]Types) {
	for key, propertie := range schema.Properties {
		t := SchemaToType(WrapContext(ctx, schema), key, propertie, schema.Required)
		fields = append(fields, t)
	}
	return fields
}

func (s *AnonymousStruct) IsNotEmpty() bool {
	return len(s.Fields) > 0
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
	{{- if isEmbeddedStruct $propertie -}}
		{{template "embeddedStruct" $propertie }}
	{{end -}}
	{{- if isReference $propertie -}}
		{{template "reference" $propertie }}
	{{end -}}
	{{- if isInterfaceReference $propertie -}}
		{{template "interfacereference" $propertie }}
	{{end -}}
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
	{{- if isInteger $propertie -}}
		{{template "integer" $propertie }}
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
