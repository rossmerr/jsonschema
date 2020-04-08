package handlers

import (
	"fmt"
	"strconv"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleOneOf(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	parent := doc.Root()

	methodSignature := parser.NewMethodSignature(name)
	types := make([]string, 0)

	for i, subschema := range schema.OneOf {
		if subschema.Ref.IsNotEmpty() {
			receiver := subschema.Ref.ToKey()
			types = append(types, receiver)
			ctx.RegisterMethodSignature(receiver, methodSignature)
			continue
		}
		structname := name + strconv.Itoa(i)
		t, err := ctx.Process(doc, structname, subschema)
		if err != nil {
			return nil, err
		}
		if _, ok := doc.Globals[structname]; !ok {
			doc.Globals[structname] = templates.NewType(structname, schema.Description, t)
		} else {
			return nil, fmt.Errorf("handleoneof: oneOf, global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		types = append(types, structname)
		ctx.RegisterMethodSignature(structname, methodSignature)

	}

	doc.AddImport("encoding/json")
	doc.Globals[name] = templates.NewInterface(name).WithMethodSignature(methodSignature)
	r := templates.NewReference(name, "", parser.NewType(name, parser.Reference), types...)

	return &templates.OneOf{r}, nil
}
