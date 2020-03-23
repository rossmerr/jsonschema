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

	interfaces, fields = addProperties(ctx, schema, interfaces, fields)
	for _, child := range schema.AllOf {
		if child.Ref != jsonschema.EmptyString {
			_, typename, _ := ResolvePointer(ctx, child.Ref)
			t := NewReference(typename)
			fields = append(fields, t)
			continue
		}

		//  if len(child.Properties) > 0 {
		// 	 //TODO fix
		// // 	interfaces, fields = addProperties(ctx, child, interfaces, fields)
		//  	continue
		//  }
		// t := SchemaToType(ctx, typename, child, schema.Required)
		// fields = append(fields, t)
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

func addProperties(ctx *SchemaContext, schema *jsonschema.Schema, interfaces []*Interface, fields []Types) ([]*Interface, []Types) {
	for key, propertie := range schema.Properties {
		t := SchemaToType(WrapContext(ctx, schema), key, propertie, schema.Required)
		if RequiesInterface(propertie) {
			if i, ok := t.(*Interface); ok {
				interfaces = append(interfaces, i)
			}
		}

		fields = append(fields, t)
	}
	return interfaces, fields
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
	{{- if isReference $propertie -}}
		{{template "reference" $propertie }}
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
