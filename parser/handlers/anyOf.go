package handlers

import (
	"fmt"
	"strconv"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleAnyOf(doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	parent := doc.Root()

	typename := parent.ID.ToTypename() + jsonschema.ToTypename(name)

	for i, subschema := range schema.AnyOf {
		if subschema.Ref.IsNotEmpty() {
			doc.AddMethods(subschema.Ref.ToTypename(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t, err := doc.Process(structname, subschema)
		if err != nil {
			return nil, err
		}
		if _, ok := doc.Globals[structname]; !ok {
			doc.Globals[structname] = types.NewRoot(subschema.Description, t, typename)
		} else {
			return nil, fmt.Errorf("handleanyof: anyOf, global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		doc.AddMethods(structname, typename)

	}

	doc.Globals[name] = types.NewInterface(typename)

	return types.NewInterfaceReference(name, "[]"+typename), nil
}
