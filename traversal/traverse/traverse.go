package traverse

import (
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/traversal"
)

// Walk, walks the Schema looking for the matching reference (case sensitive),
func Walk(s jsonschema.JsonSchema, path jsonschema.Path) *jsonschema.Schema {
	if len(path) == 0 {
		return s.(*jsonschema.Schema)
	}
	recordSchema := &recordSchema{
		match: nil,
		path:  path,
	}
	traversal.WalkSchema(s.(*jsonschema.Schema), fieldFunc, mapKeyFunc, recordSchema)
	return recordSchema.match
}

type recordSchema struct {
	match *jsonschema.Schema
	path  jsonschema.Path
}

func (s *recordSchema) Record() {}

func fieldFunc(structField reflect.StructField, field, val reflect.Value, record traversal.Record) traversal.State {

	recordSchema := record.(*recordSchema)

	if len(recordSchema.path) == 0 {
		return traversal.Return
	}

	fragment := recordSchema.path[0]

	tag := structField.Tag
	if tag == "" {
		return traversal.Continue
	}

	v, ok := tag.Lookup("json")
	if !ok {
		return traversal.Continue
	}

	tagFields := strings.Split(v, ",")
	list := jsonschema.ForEach(tagFields, func(v string) string { return strings.ToLower(v) })
	if !jsonschema.Contains(list, fragment) {
		return traversal.Continue
	}

	recordSchema.path = recordSchema.path[1:]
	if len(recordSchema.path) == 0 {
		schema := val.Interface().(*jsonschema.Schema)
		recordSchema.match = schema
		return traversal.Return
	}
	return traversal.MatchReturn
}

func mapKeyFunc(field, val reflect.Value, record traversal.Record) traversal.State {

	recordSchema := record.(*recordSchema)

	if len(recordSchema.path) == 0 {
		return traversal.Return
	}

	fragment := recordSchema.path[0]

	if field.Kind() != reflect.String {
		return traversal.Continue
	}

	if field.String() != fragment {
		return traversal.Continue
	}

	recordSchema.path = recordSchema.path[1:]
	if len(recordSchema.path) == 0 {
		if !val.IsNil() {
			i := val.Interface()
			if schema, ok := i.(*jsonschema.Schema); ok {
				recordSchema.match = schema
				return traversal.Return
			}
		}
	}
	return traversal.MatchReturn
}
