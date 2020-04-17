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
	dict[key] = emptyString

	j := tags.KeysString(dict)
	if !jsonschema.Contains(required, strings.ToLower(key)) {
		j = j + ",omitempty"
	}

	return fmt.Sprintf("json:\"%v\"", j)
}
