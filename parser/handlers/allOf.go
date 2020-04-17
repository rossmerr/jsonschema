package handlers

import (
	"fmt"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleAllOf(ctx *parser.SchemaContext, doc parser.Root, name string, schema *jsonschema.Schema) (parser.Component, error) {

	properties := map[string]*jsonschema.Schema{}

	for _, subschema := range schema.AllOf {
		switch isReference, refSchema, err := isReference(ctx, subschema, schema); {
		case err != nil:
			return nil, err
		case !isReference:
		case refSchema != nil:
			for key, prop := range refSchema.Properties {
				properties[key] = prop
			}
			continue
		case refSchema == nil:
			properties[subschema.Ref.ToKey()] = subschema
			continue
		}

		for key, prop := range subschema.Properties {
			properties[key] = prop
		}
	}

	schema.Properties = properties
	typename := strings.Trim(schema.Parent.Key+" "+name, " ")

	obj, err := HandleObject(ctx, doc, typename, schema)
	if err != nil {
		return nil, err
	}

	s, isStruct := obj.(*templates.Struct)
	if !isStruct {
		return nil, fmt.Errorf("handleallof: obj not a *templates.Struct found '%v'", obj)
	}

	doc.Globals()[typename] = templates.NewType(schema.Description, s)
	r := templates.NewReference(typename, "", parser.NewType(name, parser.Object))
	return &templates.AllOf{r}, nil
}

func isReference(ctx *parser.SchemaContext, subschema *jsonschema.Schema, schema *jsonschema.Schema) (bool, *jsonschema.Schema, error) {
	if subschema.Ref.IsNotEmpty() {
		refSchema, err := ctx.ResolveSchema(subschema.Ref, schema.Base())
		if err != nil {
			return false, nil, err
		}
		if refSchema.Type == jsonschema.Object {
			return true, refSchema, nil
		}
		return true, nil, nil
	}
	return false, nil, nil
}
