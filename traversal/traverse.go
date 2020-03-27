package traversal

import (
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

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

func walkSchema(val reflect.Value, pointer jsonschema.Pointer) (reflect.Value, bool) {
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
			arr := i.([]*jsonschema.Schema)
			for _, v := range arr {
				val := reflect.ValueOf(v).Elem()
				if schema, ok := walkSchema(val, pointer); ok {
					return schema, ok
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
