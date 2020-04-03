package handlers

import (
	"fmt"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

// Enum's always get moved up to the package level and as such always get added to the Global
// This create two different ways it can get handled, but we only have one code path for both!
//
// 1) When the enum is embedded within the Properties it will need to return a Reference as the enum will
// be moved up to the global level and the parent struct will need this reference field to the enum
//
// 2) When the enum is within the Definitions a field reference is not required, as any reference to it
// would come from a $Ref, as such no Reference is required on the return, the returned reference will get
// ignored in the calling code of HandleStruct, because it will try and add the returned Reference
// to the global level but as it's name/key will match on the already added Enum name/key bellow it will
// get ignored
func HandleEnum(ctx *parser.SchemaContext, doc *parser.Document, name string, schema *jsonschema.Schema) (parser.Types, error) {
	constItems := []*templates.ConstItem{}

	if schema.Parent != nil {
		name = strings.Join([]string{jsonschema.ToTypename(schema.Parent.Key), jsonschema.ToTypename(name)}, "_")
		name = strings.TrimLeft(name, "_")
	}

	for _, value := range schema.Enum {
		c := templates.ConstItem{
			Name:  jsonschema.ToTypename(value),
			Type:  name,
			Value: value,
		}
		constItems = append(constItems, &c)
	}
	c := templates.NewConst(constItems...)

	typenameEnum := name + "Items"
	if _, contains := doc.Globals[typenameEnum]; !contains {
		doc.Globals[typenameEnum] = c
	} else {
		return nil, fmt.Errorf("handleenum: enum, global keys need to be unique found %v more than once, in %v", name, schema.Parent.ID)
	}

	enum := templates.NewEnum(name, schema.Description, schema.Type.String(), schema.Enum, constItems)

	// The above check for the 'typenameEnum' in the global should already cover this, so no need for a second check
	doc.Globals[name] = templates.NewRoot(schema.Description, enum)

	return templates.NewReference(name, "", enum.Name()), nil
}