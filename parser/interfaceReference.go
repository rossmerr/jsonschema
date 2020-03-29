package parser

import (
	"strconv"

	"github.com/RossMerr/jsonschema"
)

type InterfaceReference struct {
	Type     string
	name     string
	FieldTag string
}

func NewInterfaceReferenceAllOf(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) Types {
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname()

	types := []string{}
	for i, subschema := range subschemas {
		if subschema.Ref.IsNotEmpty() {
			types = append(types, subschema.Ref.ToTypename())
			continue
		}
		structname := typename + strconv.Itoa(i)
		types = append(types, structname)
		t := schemaToType(ctx, NewName(structname), subschema, false)
		ctx.Globals[structname] = PrefixType(t, typename)
	}

	return NewEmbeddedStruct(name.Fieldname(), fieldTag, types...)
}

func NewInterfaceReferenceAnyOf(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) *InterfaceReference {
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname()

	for i, subschema := range subschemas {
		if subschema.Ref.IsNotEmpty() {
			ctx.AddMethods(subschema.Ref.ToTypename(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t := schemaToType(ctx, NewName(structname), subschema, false)
		ctx.Globals[structname] = PrefixType(t, typename)
		ctx.AddMethods(structname, typename)

	}

	ctx.Globals[name.Fieldname()] = NewInterface(typename)

	return &InterfaceReference{
		Type:     "[]" + typename,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}
}

func NewInterfaceReferenceOneOf(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) *InterfaceReference {
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname()

	for i, subschema := range subschemas {
		if subschema.Ref.IsNotEmpty() {
			ctx.AddMethods(subschema.Ref.ToTypename(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t := schemaToType(ctx, NewName(structname), subschema, false)
		ctx.Globals[structname] = PrefixType(t, typename)
		ctx.AddMethods(structname, typename)

	}

	ctx.Globals[name.Fieldname()] = NewInterface(typename)

	return &InterfaceReference{
		Type:     typename,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}
}

func (s *InterfaceReference) Comment() string {
	return jsonschema.EmptyString
}

func (s *InterfaceReference) Name() string {
	return s.name
}

const InterfaceReferenceTemplate = `
{{- define "interfacereference" -}}
{{ .Name}} {{ .Type}} {{ .FieldTag }}
{{- end -}}
`
