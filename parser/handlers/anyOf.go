package handlers

import (
	"fmt"
	"strconv"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleAnyOf(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	parent := doc.Root()

	typename := jsonschema.ToTypename(schema.Parent.Key + " " + name)
	methodSignature := parser.NewMethodSignature(typename)

	for i, subschema := range schema.AnyOf {
		if subschema.Ref.IsNotEmpty() {
			ctx.RegisterMethodSignature(subschema.Ref.ToTypename(), methodSignature)
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
			return nil, fmt.Errorf("handleanyof: anyOf, global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		ctx.RegisterMethodSignature(structname, methodSignature)
	}

	doc.Globals[name] = templates.NewInterface(typename).WithMethodSignature(methodSignature)
	r := templates.NewReference(name, "", "[]"+typename)
	return &templates.AnyOf{r}, nil
}
