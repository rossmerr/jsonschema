package traversal

import (
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

func Traverse(s *jsonschema.Schema, query []string) *jsonschema.Schema {
	if query == nil {
		return nil
	}

	if len(query) == 0 {
		return nil
	}

	val := reflect.ValueOf(s).Elem()
	return traverse(val, query)
}

func traverse(val reflect.Value, query []string) *jsonschema.Schema {
	if len(query) == 0 {
		i := val.Interface()
		if s, ok := i.(*jsonschema.Schema); ok {
			return s
		}
		return nil

	}

	segment := strings.ToLower(query[0])

	switch val.Kind() {
	case reflect.Struct:
		id := val.FieldByName("ID")
		text := id.String()
		if strings.ToLower(text) == segment {

			return traverse(val.Addr(), query[1:])
		}

		for i := 0; i < val.NumField(); i++ {
			tag := val.Type().Field(i).Tag
			if v, ok := tag.Lookup("json"); ok {
				tagFields := strings.Split(v, ",")
				list := jsonschema.ForEach(tagFields, func(v string) string { return strings.ToLower(v) })
				if jsonschema.Contains(list, segment) {
					return traverse(val.Field(i), query[1:])
				}
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if strings.ToLower(k.String()) == segment {
				val := val.MapIndex(k)
				return traverse(val, query[1:])
			}
		}
	case reflect.Slice:
		i := val.Interface()
		arr := i.([]*jsonschema.Schema)
		for _, v := range arr {
			val := reflect.ValueOf(v).Elem()
			result := traverse(val, query)
			if result != nil {
				return result
			}
		}
	default:
		return nil
	}
	return nil
}
