package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Struct struct {
	comment    string
	name       string
	Fields     []Types
	StructTag  string
	Globals    []Types
}

func NewStruct(ctx *SchemaContext, name *Name, schema *jsonschema.Schema, renderFieldTags bool, required ...string) *Struct {
	fields, globals := addProperties(ctx, schema, []Types{}, []Types{})

	for _, child := range append(schema.OneOf, append(schema.AnyOf, schema.AllOf...)...) {
		fields, globals = addProperties(ctx, child, fields, globals)
	}

	for _, t := range fields {
		if ref, ok := t.(*InterfaceReference); ok {
			i := NewInterface(ref.Type)
			globals = append(globals, i)
		}
	}

	var structTag string
	if renderFieldTags {
		structTag = ctx.Tags.ToFieldTag(name.Tagname(), schema, required)
	}
	return &Struct{
		comment:    schema.Description,
		name:       name.Fieldname(),
		Fields:     fields,
		Globals:globals,
		StructTag:  structTag,
	}
}

func addProperties(ctx *SchemaContext, schema *jsonschema.Schema, fields, globals []Types) ([]Types, []Types) {
	newCtx := ctx.WrapContext(schema)
	for key, propertie := range schema.Properties {
		name := NewName(key)
		fields, globals = funcName(newCtx, append(propertie.OneOf, propertie.AllOf...), name, propertie, key, fields, globals)

		if propertie.AllOf != nil && len(propertie.AllOf) > 0 {
			continue
		}

		if !propertie.Ref.IsEmpty()  {
			t := NewReference(ctx, propertie.Ref, name)
			fields = append(fields, t)
			continue
		}

		t := SchemaToType(newCtx, NewName(key), propertie, true, schema.Required...)
		fields = append(fields, t)
	}
	return fields, globals
}

func funcName(ctx *SchemaContext, arr []*jsonschema.Schema, name *Name, propertie *jsonschema.Schema, key string, fields, globals []Types) ([]Types, []Types) {
	if arr != nil && len(arr) > 0 {
		reference := NewInterfaceReference(ctx, name, propertie)
		fields = append(fields, reference)
		for _, item := range arr {
			if !item.Ref.IsEmpty() {
				continue
			}

			name := NewName(jsonschema.Structname(key + " " + item.Structname()))
			t := SchemaToType(ctx, name, item, false)
			g := PrefixType(t, reference.Type)
			globals = append(globals, g)
		}
	}
	return fields, globals
}

func (s *Struct) IsNotEmpty() bool {
	return len(s.Fields) > 0
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

{{range $key, $global := .Globals -}}
	{{- template "kind" $global -}}
{{end}}
{{end}}`
