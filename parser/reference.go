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
