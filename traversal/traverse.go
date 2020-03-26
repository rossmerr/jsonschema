package traversal

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/RossMerr/jsonschema"
)

func Traverse(s *jsonschema.Schema, pointer jsonschema.Pointer) *jsonschema.Schema {
	if pointer == nil {
		log.Print(fmt.Sprintf("Traverse: nil pointer not allowed"))
		return nil
	}

	val := reflect.ValueOf(s)
	return traverse(val, pointer)
}

func traverse(val reflect.Value, pointer jsonschema.Pointer) *jsonschema.Schema {
	if len(pointer) == 0 {
		i := val.Interface()
		if s, ok := i.(*jsonschema.Schema); ok {
			return s
		}
		return nil

	}

	segment := strings.ToLower(pointer[0])

	switch val.Kind() {
	case reflect.Struct:
		id := val.FieldByName("ID")
		text := id.String()
		if strings.ToLower(text) == segment {

			return traverse(val.Addr(), pointer[1:])
		}

		for i := 0; i < val.NumField(); i++ {
			tag := val.Type().Field(i).Tag
			if v, ok := tag.Lookup("json"); ok {
				tagFields := strings.Split(v, ",")
				list := jsonschema.ForEach(tagFields, func(v string) string { return strings.ToLower(v) })
				if jsonschema.Contains(list, segment) {
					return traverse(val.Field(i), pointer[1:])
				}
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if strings.ToLower(k.String()) == segment {
				val := val.MapIndex(k)
				return traverse(val, pointer[1:])
			}
		}
	case reflect.Slice:
		i := val.Interface()
		arr := i.([]*jsonschema.Schema)
		for _, v := range arr {
			val := reflect.ValueOf(v).Elem()
			result := traverse(val, pointer)
			if result != nil {
				return result
			}
		}
	case reflect.Ptr:
		return traverse(val.Elem(), pointer)
	default:
		return nil
	}
	return nil
}
