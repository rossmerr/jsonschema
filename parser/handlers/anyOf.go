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

	typename := parent.ID.ToTypename() + jsonschema.ToTypename(name)

	for i, subschema := range schema.AnyOf {
		if subschema.Ref.IsNotEmpty() {
			method := parser.NewMethod(subschema.Ref.ToTypename(), typename)
			ctx.AddMethods(subschema.Ref.ToTypename(), method)
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
		method := parser.NewMethod(structname, typename)
		ctx.AddMethods(structname, method)

	}

	doc.Globals[name] = templates.NewInterface(typename)

	return templates.NewInterfaceReference(name, "[]"+typename), nil
}
