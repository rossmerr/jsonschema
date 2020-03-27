package parser

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal"
)

type Reference struct {
	Type     string
	name     string
	FieldTag string
}

func NewReference(ctx *SchemaContext, ref jsonschema.Reference, name *Name, fieldTag string) *Reference {
	typename := ResolvePointer(ctx, ref)

	return &Reference{
		Type:     typename,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}
}

func (s *Reference) Comment() string {
	return jsonschema.EmptyString
}

func (s *Reference) Name() string {
	return s.name
}

func ResolvePointer(ctx *SchemaContext, ref jsonschema.Reference) string {
	path := ref.Path()
	reference := ref.Base()
	base := ctx.Parent()
	if reference != "." {
		base = ctx.References[reference]
	}

	def := traversal.Traverse(base, ref)

	fieldname := path.Fieldname()
	if fieldname == jsonschema.EmptyString {
		fieldname = jsonschema.Fieldname(def.ID.Filename())
	}

	return fieldname
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} *{{ .Type}} {{ .FieldTag }}
{{- end -}}
`
