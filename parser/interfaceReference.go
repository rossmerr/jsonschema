package parser

import "github.com/RossMerr/jsonschema"

type InterfaceReference struct {
	Type string
	name     string
	FieldTag string
}

func NewInterfaceReference(ctx *SchemaContext, name *Name, schema *jsonschema.Schema) *InterfaceReference {
	parent, _ := ctx.Base()
	fieldTag := ctx.Tags.ToFieldTag(name.Tagname(), schema, parent.Required)

	filename := parent.ID.Filename()
	typename := jsonschema.Structname(filename) + name.Fieldname()

	for _, item := range append(schema.OneOf, append(schema.AnyOf, schema.AllOf...)...) {
		structname := item.Ref.Fieldname()
		ctx.Implementations.AddMethod(structname, typename)
	}

	return &InterfaceReference{
		Type: typename,
		name: name.Fieldname(),
		FieldTag: fieldTag,
	}
}

func (s *InterfaceReference) Comment() string {
	return jsonschema.EmptyString
}

func (s *InterfaceReference) Name() string {
	return s.name
}

const InterfaceReferenceTemplate = `
{{- define "interfacereference" -}}
{{ .Name}} {{ .Type}} {{ .FieldTag }}
{{- end -}}
`