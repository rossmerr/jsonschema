package jsonschema

import (
	"reflect"
	"strings"
)

type Schema struct {
	ID                   ID             `json:"$id,omitempty"`
	Schema               string         `json:"$schema,omitempty"`
	Ref                  ID             `json:"$ref,omitempty"`
	Defs                 map[ID]*Schema `json:"$defs,omitempty"`
	Anchor               ID             `json:"$anchor,omitempty"`
	Description          string         `json:"description,omitempty"`
	Title                string         `json:"title,omitempty"`
	TypeValue            string         `json:"type,omitempty"`
	Required             []string       `json:"required,omitempty"`
	Properties           map[ID]*Schema `json:"properties,omitempty"`
	Definitions          map[ID]*Schema `json:"definitions,omitempty"`
	Items                *Schema        `json:"items,omitempty"`
	OneOf                []*Schema      `json:"oneof,omitempty"`
	AnyOf                []*Schema      `json:"anyof,omitempty"`
	AllOf                []*Schema      `json:"allof,omitempty"`

	// Validation
	MaxProperties    *uint32 `json:"maxlroperties,omitempty"`
	MinProperties    *uint32 `json:"minlroperties,omitempty"`
	MaxLength        *uint32 `json:"maxlength,omitempty"`
	MinLength        *uint32 `json:"minlength,omitempty"`
	MaxContains      *uint32 `json:"maxcontains,omitempty"`
	MinContains      *uint32 `json:"mincontains,omitempty"`
	MaxItems         *uint32 `json:"maxitems,omitempty"`
	MinItems         *uint32 `json:"minitems,omitempty"`
	Maximum          *int32  `json:"maximum,omitempty"`
	ExclusiveMaximum *int32  `json:"exclusivemaximum,omitempty"`
	Minimum          *int32  `json:"minimum,omitempty"`
	ExclusiveMinimum *int32  `json:"exclusiveminimum,omitempty"`
}

func (s Schema) Type() reflect.Kind {
	switch s.TypeValue {
	case Number:
		return reflect.Float64
	case Integer:
		return reflect.Int32
	case String:
		return reflect.String
	case Array:
		return reflect.Array
	case Object:
		if s.OneOf != nil && len(s.OneOf) > 0{
			return reflect.Interface
		}
		if s.AnyOf != nil && len(s.AnyOf) > 0 {
			return reflect.Interface
		}
		if s.AllOf != nil && len(s.AllOf) > 0 {
			return reflect.Interface
		}
		if s.Ref != EmptyString {
			return reflect.Ptr
		}
		return reflect.Struct
	}

	if s.Ref != EmptyString {
		return reflect.Ptr
	}

	return reflect.Invalid
}

func (s *Schema) Traverse(query []string) *Schema {
	if query == nil {
		return nil
	}

	if len(query) == 0 {
		return nil
	}

	val := reflect.ValueOf(s).Elem()
	return traverse(val, query)
}

func traverse(val reflect.Value, query []string) *Schema {
	if len(query) == 0 {
		i := val.Interface()
		if s, ok := i.(*Schema); ok{
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
				list := ForEach(tagFields, func(v string) string { return strings.ToLower(v) })
				if Contains(list, segment) {
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
		arr := i.([]*Schema)
		for _, v := range arr  {
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
