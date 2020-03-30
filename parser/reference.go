package parser

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal/traverse"
	"github.com/gookit/color"
)

type Reference struct {
	Type     string
	name     string
	FieldTag string
}

func NewReference(ctx *SchemaContext, ref jsonschema.Reference, name *Name, fieldTag string) (*Reference, error) {
	typename, err := ResolvePointer(ctx, ref)
	red := color.FgRed.Render

	if err != nil {
		fmt.Printf(red("🗴")+"Reference not found %v\n", ref)
	}

	return &Reference{
		Type:     typename,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}, err
}

func (s *Reference) Comment() string {
	return jsonschema.EmptyString
}

func (s *Reference) Name() string {
	return s.name
}

func ResolvePointer(ctx *SchemaContext, ref jsonschema.Reference) (string, error) {
	path := ref.Path()
	if fieldname := path.ToFieldname(); fieldname != "." {
		return fieldname, nil
	}

	base := ctx.Parent()
	if id, err := ref.ID(); err == nil {
		base = ctx.References[id]
	}

	def := traverse.Walk(base, path)
	if def == nil {
		return ".", fmt.Errorf("Path not found %v", path)
	}
	return def.ID.ToTypename(), nil
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} *{{ .Type}} {{ .FieldTag }}
{{- end -}}
`
