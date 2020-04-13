package handlers

import "github.com/RossMerr/jsonschema/parser"

func DefaultHandlers() *parser.HandlerRegistry {
	registry := parser.NewHandlerRegistry()
	registry.RegisterHandler(parser.Boolean, HandleBoolean)
	registry.RegisterHandler(parser.OneOf, HandleOneOf)
	registry.RegisterHandler(parser.AnyOf, HandleAnyOf)
	registry.RegisterHandler(parser.AllOf, HandleAllOf)
	registry.RegisterHandler(parser.Enum, HandleEnum)
	registry.RegisterHandler(parser.Array, HandleArray)
	registry.RegisterHandler(parser.Reference, HandleReference)
	registry.RegisterHandler(parser.Object, HandleObject)
	registry.RegisterHandler(parser.Number, HandleNumber)
	registry.RegisterHandler(parser.Interger, HandleInteger)
	registry.RegisterHandler(parser.String, HandleString)
	registry.RegisterHandler(parser.Document, HandleDocument)
	return registry
}
