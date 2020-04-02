package types

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/traversal/traverse"
)

var _ document.Types = (*Reference)(nil)

type Reference struct {
	Type      string
	name      string
	comment   string
	FieldTag  string
	Reference string
}

func NewReference(name, comment, typename string) *Reference {
	return &Reference{
		name:    jsonschema.ToTypename(name),
		comment: comment,
		Type:    typename,
	}
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
	return s.comment
}

func (s *Reference) Name() string {
	return s.name
}

func ResolvePointer(ctx *document.Document, ref jsonschema.Reference) (string, error) {
	path := ref.Path()
	if fieldname := path.ToFieldname(); fieldname != "." {
		return fieldname, nil
	}

	var base jsonschema.JsonSchema
	base = ctx.Root()
	if id, err := ref.ID(); err == nil {
		if err != nil {
			return ".", fmt.Errorf("resolvepointer: %w", err)

		}
		base = ctx.References[id]
	}

	def := traverse.Walk(base, path)
	if def == nil {
		return ".", fmt.Errorf("resolvepointer: path not found %v", path)
	}
	return def.ID.ToTypename(), nil
}

const ReferenceTemplate = `
{{- define "reference" -}}
{{ .Name}} {{ .Reference}}{{ .Type}} {{ .FieldTag }}
{{- end -}}
`
