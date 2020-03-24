package parser

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal"
)

func ResolvePointer(ctx *SchemaContext, ref jsonschema.Reference) (*jsonschema.Schema, string, *SchemaContext) {

	pointer := ref.Pointer()
	reference := ref.Base()
	var base *jsonschema.Schema
	if reference != "." {
		base = ctx.References[reference]
	} else {
		base, _ = ctx.Base()
	}

	def := traversal.Traverse(base, pointer)

	fieldname := pointer.Fieldname()
	if fieldname == jsonschema.EmptyString {
		fieldname = jsonschema.Fieldname(def.ID.Filename())
	}

	return def, fieldname, WrapContext(ctx, base)
}


type Reference struct {
	Type string
	Name     string
}

func NewReference(typename, name string) *Reference {
	return &Reference{
		Type: typename,
		Name:     name,
	}
}

func (s *Reference) Comment() string {
	return jsonschema.EmptyString
}

func (s *Reference) ID() string {
	return jsonschema.EmptyString
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} *{{ .Type}}
{{- end -}}
`
