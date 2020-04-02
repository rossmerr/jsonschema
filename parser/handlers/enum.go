package handlers

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleEnum(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {
	constItems := []*types.ConstItem{}
	for _, value := range schema.Enum {
		c := types.ConstItem{
			Name:  jsonschema.ToTypename(value),
			Type:  name,
			Value: value,
		}
		constItems = append(constItems, &c)
	}
	c := types.NewConst(constItems...)
	typenameEnum := name + "Enum"
	if _, ok := doc.Globals[typenameEnum]; !ok {
		doc.Globals[typenameEnum] = c
	} else {
		return nil, fmt.Errorf("handleenum: enum, global keys need to be unique found %v more than once, in %v", name, schema.Parent.ID)
	}

	return types.NewEnum(name, schema.Description, schema.Type.String(), schema.Enum, constItems), nil
}
