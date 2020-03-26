package parser

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal"
)

type Reference struct {
	Type string
	name     string
}

func NewReference(ctx *SchemaContext, ref jsonschema.Reference, name *Name) *Reference {
	_, typename, _ := ResolvePointer(ctx, ref)

	return &Reference{
		Type: typename,
		name:     name.Fieldname(),
	}
}

func (s *Reference) Comment() string {
	return jsonschema.EmptyString
}

func (s *Reference) Name() string {
	return s.name
}

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

	return def, fieldname, ctx.WrapContext(base)
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} *{{ .Type}}
{{- end -}}
`
