package traversal

import (
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

// Traverse, walks the Schema looking for the matching reference (case sensitive),
// follows the basic ideas of a JSON ID
func Traverse(s *jsonschema.Schema, path jsonschema.Path) *jsonschema.Schema {
	if len(path) == 0 {
		return s
	}
	val := reflect.ValueOf(s)
	return traverse(val, path)
}

// WalkSchema, finds all ID's and resolve there URI's
func WalkSchema(s *jsonschema.Schema) map[jsonschema.ID]*jsonschema.Schema {
	val := reflect.ValueOf(s)
	return walkSchema(val, map[jsonschema.ID]*jsonschema.Schema{})
}

// walk, search down the entire schema looking for the first matching ID field,
// a ID field can be on any type but if the type does not end on a Schema then any Path should!
func walkSchema(val reflect.Value, uri map[jsonschema.ID]*jsonschema.Schema) map[jsonschema.ID]*jsonschema.Schema {
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Kind() == reflect.String {
				if field.Type().Name() == "ID" {
					id := jsonschema.NewID(field.String())
					if id != "." {
						i := val.Addr().Interface()
						uri[id] = i.(*jsonschema.Schema)
					}
				}
				continue
			}
			return walkSchema(field, uri)
		}
	case reflect.Map:
		for _, v := range val.MapKeys() {
			val := val.MapIndex(v)
			uri = walkSchema(val, uri)
		}
	case reflect.Slice:
		if !val.IsNil() {
			i := val.Interface()
			if arr, ok := i.([]*jsonschema.Schema); ok {
				for _, v := range arr {
					val := reflect.ValueOf(v).Elem()
					uri = walkSchema(val, uri)
				}
			}
		}
	case reflect.Ptr:
		uri = walkSchema(val.Elem(), uri)
	default:
		return uri
	}
	return uri
}

// traverse, from the val search down each fragment of the path until all parts are matched.
// Path fragment must match there json tag names to revolve JSON Pointers and end on a schema.
func traverse(val reflect.Value, path jsonschema.Path) *jsonschema.Schema {
	if len(path) == 0 {
		i := val.Interface()
		if s, ok := i.(*jsonschema.Schema); ok {
			return s
		}
		return nil

	}

	segment := path[0]

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			tag := val.Type().Field(i).Tag
			if v, ok := tag.Lookup("json"); ok {
				tagFields := strings.Split(v, ",")
				list := jsonschema.ForEach(tagFields, func(v string) string { return strings.ToLower(v) })
				if jsonschema.Contains(list, segment) {
					return traverse(val.Field(i), path[1:])
				}
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if k.String() == segment {
				val := val.MapIndex(k)
				return traverse(val, path[1:])
			}
		}
	case reflect.Slice:
		i := val.Interface()
		arr := i.([]*jsonschema.Schema)
		for _, v := range arr {
			val := reflect.ValueOf(v).Elem()
			result := traverse(val, path)
			if result != nil {
				return result
			}
		}
	case reflect.Ptr:
		return traverse(val.Elem(), path)
	default:
		return nil
	}
	return nil
}
