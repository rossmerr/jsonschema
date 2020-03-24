package parser

import "github.com/RossMerr/jsonschema"

type InterfaceReference struct {
	Type string
	Name     string
}


func NewInterfaceReference(ctx *SchemaContext,  name string, schema *jsonschema.Schema) *InterfaceReference {
	parent, _ := ctx.Base()
	filename := parent.ID.Filename()
	typename := jsonschema.Structname(filename) + name
	var arr []*jsonschema.Schema
	if schema.OneOf != nil {
		arr = schema.OneOf
	}
	if schema.AnyOf != nil {
		arr = schema.AnyOf
	}

	for _, item := range arr {
		fieldname := item.Ref.Fieldname()
		if fieldname != jsonschema.EmptyString {
			arr := ctx.Implementations[fieldname]
			if arr == nil {
				arr = []string{}
			}
			arr = append(arr, typename)
			ctx.Implementations[fieldname] = arr
		}
	}

	return &InterfaceReference{
		Type: typename,
		Name:     name,
	}
}

func (s *InterfaceReference) Comment() string {
	return jsonschema.EmptyString
}

func (s *InterfaceReference) ID() string {
	return jsonschema.EmptyString
}

const InterfaceReferenceTemplate = `
{{- define "interfacereference" -}}
{{ .Name}} {{ .Type}}
{{- end -}}
`

