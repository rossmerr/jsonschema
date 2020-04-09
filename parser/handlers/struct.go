package handlers

import (
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/tags"
	"github.com/RossMerr/jsonschema/parser/tags/json"
	"github.com/RossMerr/jsonschema/parser/tags/validate"
	"github.com/RossMerr/jsonschema/parser/templates"
)

var fieldTags = tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})

func HandleObject(ctx *parser.SchemaContext, doc parser.Root, name string, schema *jsonschema.Schema) (parser.Component, error) {

	fields := []parser.Component{}
	for key, propertie := range schema.Properties {
		s, err := ctx.Process(doc, key, propertie)
		if err != nil {
			return nil, err
		}

		if f, ok := s.(parser.Field); ok {
			fieldTag := fieldTags.ToFieldTag(key, propertie, schema.Required)
			ref := !jsonschema.Contains(schema.Required, strings.ToLower(key))
			f.WithFieldTag(fieldTag).WithReference(ref)
		}

		fields = append(fields, s)
	}

	for key, def := range schema.AllDefinitions() {
		t, err := ctx.Process(doc, key, def)
		if err != nil {
			return nil, err
		}

		if _, ok := t.(*templates.OneOf); !ok {
			t = templates.NewType(schema.Description, t)

			if _, contains := doc.Globals()[key]; !contains {
				doc.Globals()[key] = t
			}
		}
	}

	return templates.NewStruct(name, schema.Description, fields...), nil
}
