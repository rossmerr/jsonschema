package parser

import (
	"fmt"
	"strconv"

	"github.com/RossMerr/jsonschema"
)

type InterfaceReference struct {
	Type     string
	name     string
	FieldTag string
}

func NewInterfaceReferenceAllOf(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) (Types, error) {
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
		t, err := schemaToType(ctx, NewName(structname), subschema, false)
		if err != nil {
			return nil, err
		}
		if _, ok := ctx.Globals[structname]; !ok {
			ctx.Globals[structname] = PrefixType(t, typename)
		} else {
			return nil, fmt.Errorf("Global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
	}

	return NewEmbeddedStruct(name.Fieldname(), fieldTag, types...), nil
}

func NewInterfaceReferenceAnyOf(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) (*InterfaceReference, error) {
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname()

	for i, subschema := range subschemas {
		if subschema.Ref.IsNotEmpty() {
			ctx.AddMethods(subschema.Ref.ToTypename(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t, err := schemaToType(ctx, NewName(structname), subschema, false)
		if err != nil {
			return nil, err
		}
		if _, ok := ctx.Globals[structname]; !ok {
			ctx.Globals[structname] = PrefixType(t, typename)
		} else {
			return nil, fmt.Errorf("Global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		ctx.AddMethods(structname, typename)

	}

	ctx.Globals[name.Fieldname()] = NewInterface(typename)

	return &InterfaceReference{
		Type:     "[]" + typename,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}, nil
}

func NewInterfaceReferenceOneOf(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) (*InterfaceReference, error) {
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname()

	for i, subschema := range subschemas {
		if subschema.Ref.IsNotEmpty() {
			ctx.AddMethods(subschema.Ref.ToTypename(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t, err := schemaToType(ctx, NewName(structname), subschema, false)
		if err != nil {
			return nil, err
		}
		if _, ok := ctx.Globals[structname]; !ok {
			ctx.Globals[structname] = PrefixType(t, typename)
		} else {
			return nil, fmt.Errorf("Global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		ctx.AddMethods(structname, typename)

	}

	ctx.Globals[name.Fieldname()] = NewInterface(typename)

	return &InterfaceReference{
		Type:     typename,
		name:     name.Fieldname(),
		FieldTag: fieldTag,
	}, nil
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
