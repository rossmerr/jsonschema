package jsonschema

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Schema struct {
	ID          ID                 `json:"$id,omitempty"`
	Schema      string             `json:"$schema,omitempty"`
	Ref         Pointer            `json:"$ref,omitempty"`
	Defs        map[string]*Schema `json:"$defs,omitempty"`
	Anchor      Anchor             `json:"$anchor,omitempty"`
	Description string             `json:"description,omitempty"`
	Title       string             `json:"title,omitempty"`
	Type        Kind               `json:"type,omitempty"`
	Required    []string           `json:"required,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	// Deprecated use Defs
	Definitions          map[string]*Schema `json:"definitions,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	OneOf                []*Schema          `json:"oneof,omitempty"`
	AnyOf                []*Schema          `json:"anyof,omitempty"`
	AllOf                map[string]*Schema          `json:"-"`
	Enum                 []string           `json:"enum,omitempty"`
	AdditionalProperties *bool              `json:"additionalproperties,omitempty"`

	// Validation
	MaxProperties    *uint32 `json:"maxproperties,omitempty"`
	MinProperties    *uint32 `json:"minproperties,omitempty"`
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
	Pattern          string  `json:"pattern,omitempty"`
}

func (s *Schema) IsEnum() bool {
	return s.Enum != nil
}

func (s *Schema) UnmarshalJSON(b []byte) (err error) {
	type Alias Schema
	a := Alias{}

	if err = json.Unmarshal(b, &a); err == nil {
		*s = Schema(a)
	}

	m := make(map[string]json.RawMessage)

	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	for k, v := range m {
		delete(m, k)
		m[strings.ToLower(k)] = v
	}

	s.AllOf = s.combining("allof", JsonTags(s), m)
	return
}


func (s *Schema) combining(key string, jsonTags []string, m map[string]json.RawMessage) map[string]*Schema  {
	combined := map[string]*Schema{}
	if raw, ok := m[key]; ok {
		allOf := make([]json.RawMessage, 0)
		json.Unmarshal(raw, &allOf)
		for _, k := range allOf {
			items := make(map[string]json.RawMessage)
			json.Unmarshal(k, &items)

			for key, v := range items {
				if Contains(jsonTags, key) {
					ref := Schema{}
					if err := json.Unmarshal(k, &ref); err == nil {
						combined[key] = &ref
					}
				} else {
					obj := Schema{}
					if err := json.Unmarshal(v, &obj); err == nil {
						combined[key] = &obj
					}
				}
			}
		}
	}

	return combined
}

func JsonTags(s *Schema) []string {
	tags := make([]string, 0)
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag
		if v, ok := tag.Lookup("json"); ok {
			tagFields := strings.Split(v, ",")
			list := ForEach(tagFields, func(v string) string { return strings.ToLower(v) })
			tags = append(tags, list...)
		}
	}

	return tags
}
