package types

import (
	"fmt"
	"strconv"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

func HandleOneOf(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name

	for i, subschema := range schema.OneOf {
		if subschema.Ref.IsNotEmpty() {
			ctx.AddMethods(subschema.Ref.ToTypename(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t, err := ctx.Process(structname, subschema)
		if err != nil {
			return nil, err
		}
		if _, ok := ctx.Globals[structname]; !ok {
			ctx.Globals[structname] = PrefixType(t, typename)
		} else {
			return nil, fmt.Errorf("oneOf, global keys need to be unique found %v more than once, in %v", structname, parent.ID)
		}
		ctx.AddMethods(structname, typename)

	}

	ctx.Globals[name] = NewInterface(typename)

	return &InterfaceReference{
		Type: typename,
		name: jsonschema.ToTypename(name),
	}, nil
}
