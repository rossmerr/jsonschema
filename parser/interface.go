package parser

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type Interface struct {
	id jsonschema.ID
	Package    string
	comment string
	CommentImplementations string
	Name    string
	Method string
	StructTag  string
	Structs []*AnonymousStruct
}

func NewInterface(ctx *SchemaContext, schema, parent *jsonschema.Schema) *Interface {
	name := schema.ID.Typename()

	method := MixedCase(name)

	structTag := ""
	if parent != nil {
		structTag = ctx.Tags.ToFieldTag(strings.Title(name), schema, parent)
	}

	refs := []string{}
	structs := []*AnonymousStruct{}
	if schema.OneOf != nil {
		for _, oneOf:= range schema.OneOf  {
			_, query, _ := oneOf.Ref.Pointer()
			def := parent.Traverse(query)
			def.ID = jsonschema.ID(query[len(query) -1])
			t := SchemaToType(ctx, schema.ID, def, schema)
			i := t.(*AnonymousStruct)
			i.Method = method
			structs = append(structs, i)
			refs = append(refs, oneOf.Ref.String())
		}
	}

	return &Interface{
		id: schema.ID,
		comment: schema.Description,
		CommentImplementations:strings.Join(refs, "\n // "),
		Name:    name,
		Method: method,
		Package:    ctx.Package,
		StructTag:structTag,
		Structs: structs,
	}
}

func (s *Interface) Comment() string {
	return s.comment
}

func (s *Interface) ID() jsonschema.ID {
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
{{end}}

{{end}}
`
