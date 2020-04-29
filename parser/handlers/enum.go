package handlers

import (
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
	log "github.com/sirupsen/logrus"
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
func HandleEnum(ctx *parser.SchemaContext, doc parser.Root, name string, schema *jsonschema.Schema) (parser.Component, error) {
	constItems := []*templates.ConstItem{}

	name = strings.Trim(schema.Parent.Key+" "+name, " ")
	for _, value := range schema.Enum {
		c := templates.ConstItem{
			Name:  value,
			Type:  name,
			Value: value,
		}
		constItems = append(constItems, &c)
	}
	c := templates.NewConst(constItems...)

	typenameEnum := name + " Items"
	enum := templates.NewEnum(name, schema.Description, schema.Type.String(), schema.Enum, constItems)

	if _, contains := doc.Globals()[typenameEnum]; !contains {
		doc.Globals()[typenameEnum] = c
		doc.Globals()[name] = enum
	} else {
		log.Infof("handleenum: enum, already found global key %v ignoring", name)
	}

	doc.AddImport("encoding/json")
	doc.AddImport("fmt")
	return templates.NewReference(name, "", parser.NewType(enum.Name(), parser.Enum)), nil
}
