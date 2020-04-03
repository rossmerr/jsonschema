package json

import (
	"fmt"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/tags"
)

type json struct {
}

func NewJSONTags() tags.StructTag {
	return &json{}
}

func (s *json) ToStructTag(key string, schema *jsonschema.Schema, required []string) string {

	dict := map[string]string{}

	dict[key] = jsonschema.EmptyString
	if !jsonschema.Contains(required, strings.ToLower(key)) {
		dict["omitempty"] = jsonschema.EmptyString
	}

	if len(dict) == 0 {
		return jsonschema.EmptyString
	}

	return fmt.Sprintf("json:\"%v\"", jsonschema.KeysString(dict))
}