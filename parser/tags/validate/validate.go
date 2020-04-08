package validate

import (
	"fmt"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/tags"
)

type validate struct {
}

func NewValidateTags() tags.StructTag {
	return &validate{}
}

func (s *validate) ToStructTag(key string, schema *jsonschema.Schema, required []string) string {

	dict := map[string]string{}

	if jsonschema.Contains(required, strings.ToLower(key)) {
		dict["required"] = EmptyString
	}

	if schema != nil {
		if schema.AnyOf != nil {
			dict["anyof"] = EmptyString
		}

		if schema.AllOf != nil {
			dict["allof"] = EmptyString
		}

		if schema.OneOf != nil {
			dict["oneof"] = EmptyString
		}

		if schema.MaxLength != nil {
			dict["max"] = fmt.Sprint(*schema.MaxLength)
		}

		if schema.MinLength != nil {
			dict["min"] = fmt.Sprint(*schema.MinLength)
		}

		if schema.Maximum != nil {
			dict["lte"] = fmt.Sprint(*schema.Maximum)
		}

		if schema.ExclusiveMaximum != nil {
			dict["lt"] = fmt.Sprint(*schema.ExclusiveMaximum)
		}

		if schema.Minimum != nil {
			dict["gte"] = fmt.Sprint(*schema.Minimum)
		}

		if schema.ExclusiveMinimum != nil {
			dict["gt"] = fmt.Sprint(*schema.ExclusiveMinimum)
		}

		if schema.Pattern != EmptyString {
			dict["regex"] = fmt.Sprint(schema.Pattern)
		}
	}

	if len(dict) == 0 {
		return EmptyString
	}

	return fmt.Sprintf("validate:\"%v\"", tags.KeysString(dict))
}
