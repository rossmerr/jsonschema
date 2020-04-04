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

	for i, subschema := range schema.OneOf {
		if subschema.Ref.IsNotEmpty() {
			ctx.ImplementMethodSignature(subschema.Ref.ToKey(), methodSignature)
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
		ctx.ImplementMethodSignature(structname, methodSignature)

	}

	doc.Globals[name] = templates.NewInterface(typename)
	r := templates.NewReference(name, "", typename)

	return &templates.OneOf{r}, nil
}
