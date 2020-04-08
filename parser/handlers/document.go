package handlers

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleDocument(ctx *parser.SchemaContext, doc parser.Root, name string, schema *jsonschema.Schema) (parser.Component, error) {
	document := templates.NewDocument(schema)
	for key, propertie := range schema.Properties {
		t, err := ctx.Process(document, key, propertie)
		if err != nil {
			return nil, err
		}

		switch t.(type) {
		case *templates.OneOf, *templates.AnyOf:
			continue
		default:
			t = templates.NewType(schema.Description, t)
			if _, contains := document.Globals()[key]; !contains {
				document.Globals()[key] = t
			}
		}

	}
	for key, def := range schema.AllDefinitions() {
		t, err := ctx.Process(document, key, def)
		if err != nil {
			return nil, err
		}
		if _, ok := t.(*templates.OneOf); !ok {
			t = templates.NewType(schema.Description, t)

			if _, contains := document.Globals()[key]; !contains {
				document.Globals()[key] = t
			}
		}
	}

	return document, nil
}
