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

	for i, subschema := range schema.OneOf {
		if subschema.Ref.IsNotEmpty() {
			method := parser.NewMethod(subschema.Ref.ToKey(), typename)
			ctx.ImplementInterface(subschema.Ref.ToKey(), method)
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
		method := parser.NewMethod(structname, typename)
		ctx.ImplementInterface(structname, method)

	}

	doc.Globals[name] = templates.NewInterface(typename)

	t:=  templates.NewInterfaceReference(name, typename)
	return &templates.OneOf{t}, nil
}