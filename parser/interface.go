package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Interface struct {
	id string
	Package    string
	comment string
	CommentImplementations string
	Name    string
	Method string
	StructTag  string
	Structs []*AnonymousStruct
}

func NewInterface(ctx *SchemaContext, typename string, schema *jsonschema.Schema, required []string) *Interface {
	method := MixedCase(typename)

	refs := []string{}
	structs := []*AnonymousStruct{}

	if schema.OneOf != nil {
		for _, oneOf:= range schema.OneOf  {
			fragments := oneOf.Ref.Fragments()
			def := ctx.ParentSchema.Traverse(fragments)
			typename := fragments[len(fragments)-1]
			t := SchemaToType(WrapContext(ctx, schema), typename, def, schema.Required)
			if i, ok := t.(*AnonymousStruct); ok {
				i.Method = method
				structs = append(structs, i)
				refs = append(refs, oneOf.Ref.String())
			}
		}
	}

	return &Interface{
		id: schema.ID.String(),
		comment: schema.Description,
		CommentImplementations:strings.Join(refs, "\n // "),
		Name:    typename,
		Method: method,
		Package:    ctx.Package,
		StructTag: ctx.Tags.ToFieldTag(typename, schema, required),
		Structs: structs,
	}
}

func (s *Interface) Comment() string {
	return s.comment
}

func (s *Interface) ID() string {
	return s.id
}

const InterfaceTemplate = `
{{- define "interface" -}}

{{if .Comment -}}
// {{ .Comment}}
{{ else -}}
// {{ .Name }}
{{end -}}
{{if .CommentImplementations -}}
// {{ .CommentImplementations}}
{{end -}}
type {{ .Name }} interface {
	{{ .Method}}()
}

{{range $key, $struct := .Structs -}}
	 type {{template "struct" $struct }}
{{end -}}

{{end -}}
`
