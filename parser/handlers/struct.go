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

func HandleObject(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {

	fields := []parser.Types{}
	for key, propertie := range schema.Properties {
		s, err := ctx.Process(key, propertie)
		if err != nil {
			return nil, err
		}

		tags := tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})
		fieldTag := tags.ToFieldTag(key, propertie, schema.Required)

		ref := !jsonschema.Contains(schema.Required, strings.ToLower(key))

		s.WithFieldTag(fieldTag).WithReference(ref)

		if _, ok := s.(*templates.Root); ok {
			continue
		}

		fields = append(fields, s)
	}

	for key, def := range schema.AllDefinitions() {
		t, err := ctx.Process(key, def)
		if err != nil {
			return nil, err
		}

		if _, contains := doc.Globals[key]; !contains {
			doc.Globals[key] = templates.NewRoot(schema.Description, t)
		}
	}

	return templates.NewStruct(name, schema.Description, fields), nil
}