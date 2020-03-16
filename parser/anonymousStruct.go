package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type AnonymousStruct struct {
	comment     string
	Name        string
	id          jsonschema.ID
	Fields      []Types
	StructTag   string
	Definitions []Types
	InterfaceMethods []string

}

func NewAnonymousStruct(ctx *SchemaContext, id jsonschema.ID, schema, parent *jsonschema.Schema) *AnonymousStruct {
	fields := []Types{}

	for key, propertie := range schema.Properties {
		t := SchemaToType(ctx, key, propertie, schema)

		fields = append(fields, t)
	}

	name := schema.ID.Typename()
	if name == "" {
		name = id.Typename()
	}

	structTag := ""
	if parent != nil {
		structTag = ctx.Tags.ToFieldTag(strings.Title(name), schema, parent)
	}


	definitions := []Types{}
	for key, definition := range schema.Definitions {
		t := SchemaToType(ctx, jsonschema.NewDefinitionsID(key), definition, nil)

		definitions = append(definitions, t)
	}

	return &AnonymousStruct{
		comment:     schema.Description,
		Name:        strings.Title(name),
		id:          id,
		Fields:      fields,
		StructTag:   structTag,
		Definitions: definitions,
		InterfaceMethods:jsonschema.Unique(ctx.ImplementInterface[id]),

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
		{{ $propertie.Name }} interface{} {{ $propertie.StructTag }}
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
} {{ .StructTag }}

{{ if .InterfaceMethods }}
	{{range $key, $mathod := .InterfaceMethods -}}
	func (s *{{$.Name}}) {{ $mathod }}() {}
	{{end -}}
{{end -}}
{{end}}`
