package tags

import (
	"strings"

	"github.com/RossMerr/jsonschema"
)

type FieldTag interface {
	ToFieldTag(key string, schema, parent *jsonschema.Schema) string
}

type fieldTag struct {
	structTags []StructTag
}

func NewFieldTag(structTags []StructTag) FieldTag {
	return &fieldTag{
		structTags: structTags,
	}
}

func (s *fieldTag) ToFieldTag(key string, schema, parent *jsonschema.Schema) string {
	if len(s.structTags) == 0 {
		return jsonschema.EmptyString
	}

	fieldTags := []string{}
	for _, tag := range s.structTags {
		tag := tag.ToStructTag(key, schema, parent)
		if tag != jsonschema.EmptyString {
			fieldTags = append(fieldTags, tag)
		}
	}

	return "`" + strings.Join(fieldTags, ", ") + "`"
}
