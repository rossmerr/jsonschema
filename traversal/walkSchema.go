package traversal

import (
	"reflect"

	"github.com/RossMerr/jsonschema"
)


type FieldFunc func(structField reflect.StructField, field, val reflect.Value, record Record) State

type MapKeyFunc func(field, val reflect.Value, record Record) State

// WalkSchema, search down the entire schema looking for matching field or mapKey,
func WalkSchema(s *jsonschema.Schema, fieldFunc FieldFunc, mapKeyFunc MapKeyFunc, record Record) {
	val := reflect.ValueOf(s)
	walkSchema(val, fieldFunc, mapKeyFunc, record)

}

func walkSchema(val reflect.Value, fieldFunc FieldFunc, mapKeyFunc MapKeyFunc, record Record) {
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			switch state := fieldFunc(val.Type().Field(i), field, val, record); state {
			case Match:
				walkSchema(field, fieldFunc, mapKeyFunc, record)
			case MatchReturn:
				walkSchema(val.Field(i), fieldFunc, mapKeyFunc, record)
				return
			case Return:
				return
			case Continue:

			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			val := val.MapIndex(k)
			switch state := mapKeyFunc(k, val, record); state {
			case Match:
				walkSchema(val, fieldFunc, mapKeyFunc, record)
			case MatchReturn:
				walkSchema(val, fieldFunc, mapKeyFunc, record)
				return
			case Return:
				return
			case Continue:

			}

		}
	case reflect.Slice:
		if !val.IsNil() {
			i := val.Interface()
			if arr, ok := i.([]*jsonschema.Schema); ok {
				for _, v := range arr {
					val := reflect.ValueOf(v).Elem()
					walkSchema(val, fieldFunc, mapKeyFunc, record)
				}
			}
		}
	case reflect.Ptr:
		walkSchema(val.Elem(), fieldFunc, mapKeyFunc, record)
	}
}
