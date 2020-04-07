package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleOneOf(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	parent := doc.Root()

	typename := name
	if schema.Parent != nil {
		typename = strings.Join([]string{jsonschema.ToTypename(schema.Parent.Key), jsonschema.ToTypename(name)}, "_")
		typename = strings.TrimLeft(typename, "_")
	}

	methodSignature := parser.NewMethodSignature(typename)
	types := make([]string, 0)

	for i, subschema := range schema.OneOf {
		if subschema.Ref.IsNotEmpty() {
			receiver := subschema.Ref.ToKey()
			types = append(types, jsonschema.ToTypename(receiver))
			ctx.RegisterMethodSignature(receiver, methodSignature)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t, err := ctx.Process(structname, subschema)
		if err != nil {
			return nil, err
		}
		if _, ok := doc.Globals[structname]; !ok {
			doc.Globals[structname] = templates.NewRoot(subschema.Description, t)
		} else {
			return nil, fmt.Errorf("handleoneof: oneOf, global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		types = append(types, jsonschema.ToTypename(structname))
		ctx.RegisterMethodSignature(structname, methodSignature)

	}

	doc.AddImport("encoding/json")
	doc.Globals[name] = templates.NewInterface(typename).WithMethodSignature(methodSignature)
	r := templates.NewReference(name, "", parser.NewType(typename, parser.Object), types...)

	return &templates.OneOf{r}, nil
}
