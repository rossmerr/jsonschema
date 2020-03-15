package json

import (
	"fmt"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/tags"
)

type json struct {
}

func NewJSONTags() tags.StructTag {
	return &json{}
}

func (s *json) ToStructTag(key string, schema, parent *jsonschema.Schema) string {

	dict := map[string]string{}

	dict[key] = jsonschema.EmptyString
	if parent != nil {
		if jsonschema.Contains(parent.Required, key) {
			dict["omitempty"] = jsonschema.EmptyString
		}
	}

	if len(dict) == 0 {
		return jsonschema.EmptyString
	}

	return fmt.Sprintf("json:\"%v\"", jsonschema.KeysString(dict))
}
