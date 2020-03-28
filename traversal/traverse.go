package traversal

import (
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

// Traverse, walks the Schema looking for the matching reference,
// follows the basic ideas of a JSON Pointer
func Traverse(s *jsonschema.Schema, reference jsonschema.Reference) *jsonschema.Schema {
	if reference == jsonschema.EmptyString {
		return nil
	}

	val := reflect.ValueOf(s)
	pointer, path := reference.Stat()
	if schema, ok := walkSchema(val, pointer); ok {
		return traverse(schema, path)
	}
	return traverse(val, path)
}

// walkSchema, search down the entire schema looking for the first matching ID field,
// a ID field can be on any type but if the type does not end on a Schema then any Path should!
func walkSchema(val reflect.Value, pointer jsonschema.Pointer) (reflect.Value, bool) {
	if pointer == "" {
		return val, false
	}
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Kind() == reflect.String {
				if field.Type().Name() == "ID" {
					id := field.String()
					if id == pointer.String() {
						return val.Addr(), true
					}
				}
			}
			if schema, ok := walkSchema(field, pointer); ok {
				return schema, ok
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			val := val.MapIndex(k)
			if schema, ok := walkSchema(val, pointer); ok {
				return schema, ok
			}
		}
	case reflect.Slice:
		if !val.IsNil() {
			i := val.Interface()
			if arr, ok := i.([]*jsonschema.Schema); ok {
				for _, v := range arr {
					val := reflect.ValueOf(v).Elem()
					if schema, ok := walkSchema(val, pointer); ok {
						return schema, ok
					}
				}
			}

		}
	case reflect.Ptr:
		return walkSchema(val.Elem(), pointer)
	default:
		return reflect.Value{}, false
	}
	return reflect.Value{}, false
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

	segment := strings.ToLower(path[0])

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
			if strings.ToLower(k.String()) == segment {
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
