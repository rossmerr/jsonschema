package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
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

	return HandleObject(ctx, doc, name, schema)
}
