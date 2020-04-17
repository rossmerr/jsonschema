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

		if subschema.Ref.IsNotEmpty() {
			properties[subschema.Ref.ToKey()] = subschema
			continue
		}
		for key, prop := range subschema.Properties {
			properties[key] = prop
		}
	}

	schema.Properties = properties
	typename :=  strings.Trim(schema.Parent.Key + " " + name, " ")

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
