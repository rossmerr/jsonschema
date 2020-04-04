package handlers

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleAllOf(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	properties := map[string]*jsonschema.Schema{}

	for _, subschema := range schema.AllOf {

		if subschema.Ref.IsNotEmpty() {
			properties[subschema.Ref.ToTypename()] = subschema
			continue
		}
		for key, prop := range subschema.Properties {
			properties[key] = prop
		}
	}

	schema.Properties = properties

	obj, err := HandleObject(ctx, doc, name, schema)
	if err != nil {
		return nil, err
	}

	s, ok := obj.(*templates.Struct)
	if !ok {
		return nil, fmt.Errorf("handleallof: obj not a *templates.Struct found '%v'", obj)
	}

	return &templates.AllOf{s}, nil
}
