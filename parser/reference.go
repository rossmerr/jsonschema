package parser

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal"
)

func ResolvePointer(ctx *SchemaContext, ref jsonschema.Pointer) (*jsonschema.Schema, string, *SchemaContext) {

	fragments := ref.Fragments()
	reference := ref.Base()
	var base *jsonschema.Schema
	if reference != "." {
		base = ctx.References[reference]
	} else {
		base, _ = ctx.Base()
	}

	def := traversal.Traverse(base, fragments)

	typename := fragments[len(fragments)-1]

	return def, typename, WrapContext(ctx, base)
}


type Reference struct {
	Name     string
}

func NewReference(typename string) *Reference {
	return &Reference{
		Name:     typename,
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
*{{ .Name}}
{{- end -}}
`
