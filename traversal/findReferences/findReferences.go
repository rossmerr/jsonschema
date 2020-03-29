package findReferences

import (
	"reflect"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal"
)

// Walk, finds all ID's and resolve there URI's
func Walk(s *jsonschema.Schema) map[jsonschema.ID]*jsonschema.Schema {
	recordMap := &recordMap{uri: map[jsonschema.ID]*jsonschema.Schema{}}
	traversal.WalkSchema(s, fieldFunc, mapKeyFunc, recordMap)
	return recordMap.uri
}

type recordMap struct {
	uri map[jsonschema.ID]*jsonschema.Schema
}

func (s *recordMap) Record() {}

func mapKeyFunc(field, val reflect.Value, record traversal.Record) traversal.State {
	return traversal.Match
}

func fieldFunc(structField reflect.StructField, field, val reflect.Value, record traversal.Record) traversal.State {
	recordMap := record.(*recordMap)

	if field.Kind() != reflect.String {
		return traversal.Match
	}

	if field.Type().Name() != "ID" {
		return traversal.Match
	}

	if id := jsonschema.NewID(field.String()); id != "." {
		i := val.Addr().Interface()
		recordMap.uri[id] = i.(*jsonschema.Schema)
		return traversal.Continue
	}

	return traversal.Match
}
