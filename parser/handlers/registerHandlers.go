package handlers

import "github.com/RossMerr/jsonschema/parser"

func DefaultHandlers(p parser.Parser) parser.Parser {
	p.HandlerFunc(parser.Boolean, HandleBoolean)
	p.HandlerFunc(parser.OneOf, HandleOneOf)
	p.HandlerFunc(parser.AnyOf, HandleAnyOf)
	p.HandlerFunc(parser.AllOf, HandleAllOf)
	p.HandlerFunc(parser.Enum, HandleEnum)
	p.HandlerFunc(parser.Array, HandleArray)
	p.HandlerFunc(parser.Reference, HandleReference)
	p.HandlerFunc(parser.Object, HandleObject)
	p.HandlerFunc(parser.Number, HandleNumber)
	p.HandlerFunc(parser.Interger, HandleInteger)
	p.HandlerFunc(parser.String, HandleString)
	p.HandlerFunc(parser.Root, HandleDocument)
	return p
}
