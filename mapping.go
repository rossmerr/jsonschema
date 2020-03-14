package jsonschema

type Mapping struct {
	ID string
	Document    *Schema
	Definitions *SchemaReferences
	Config      Config
}
