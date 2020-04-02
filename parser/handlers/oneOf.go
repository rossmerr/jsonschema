package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleOneOf(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	parent := doc.Root()

	typename := name
	if schema.Parent != nil {
		typename = strings.Join([]string{jsonschema.ToTypename(schema.Parent.Key), jsonschema.ToTypename(name)}, "_")
		typename = strings.TrimLeft(typename, "_")
	}

	for i, subschema := range schema.OneOf {
		if subschema.Ref.IsNotEmpty() {
			doc.AddMethods(subschema.Ref.ToKey(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t, err := doc.Process(structname, subschema)
		if err != nil {
			return nil, err
		}
		if _, ok := doc.Globals[structname]; !ok {
			doc.Globals[structname] = types.NewRoot(subschema.Description, t)
		} else {
			return nil, fmt.Errorf("handleoneof: oneOf, global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		doc.AddMethods(structname, typename)

	}

	doc.Globals[name] = types.NewInterface(typename)

	return types.NewInterfaceReference(name, typename), nil
}
