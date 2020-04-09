package parser

import "github.com/RossMerr/jsonschema"

type Resolve func(schema *jsonschema.Schema, document Root) HandleSchemaFunc
type Process func(schema *jsonschema.Schema) (Component, error)
type HandleSchemaFunc func(*SchemaContext, Root, string, *jsonschema.Schema) (Component, error)
