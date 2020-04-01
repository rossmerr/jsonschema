package types

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/traversal/traverse"
	"github.com/gookit/color"
)

var _ document.Types = (*Reference)(nil)

type Reference struct {
	Type      string
	name      string
	FieldTag  string
	Reference string
}

func HandleReference(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	typename, err := ResolvePointer(ctx, schema.Ref)
	red := color.FgRed.Render

	if err != nil {
		fmt.Printf(red("🗴")+"reference not found %v\n", schema.Ref)
	}

	return &Reference{
		Type: typename,
		name: jsonschema.ToTypename(name),
	}, err
}

func (s *Reference) WithReference(ref bool) document.Types {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Reference) WithFieldTag(tags string) document.Types {
	s.FieldTag = tags
	return s
}
func (s *Reference) Comment() string {
	return jsonschema.EmptyString
}

func (s *Reference) Name() string {
	return s.name
}

func ResolvePointer(ctx *document.DocumentContext, ref jsonschema.Reference) (string, error) {
	path := ref.Path()
	if fieldname := path.ToFieldname(); fieldname != "." {
		return fieldname, nil
	}

	var base jsonschema.JsonSchema
	base = ctx.Parent()
	if id, err := ref.ID(); err == nil {
		base = ctx.References[id]
	}

	def := traverse.Walk(base, path)
	if def == nil {
		return ".", fmt.Errorf("path not found %v", path)
	}
	return def.ID.ToTypename(), nil
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} {{ .Reference}}{{ .Type}} {{ .FieldTag }}
{{- end -}}
`
