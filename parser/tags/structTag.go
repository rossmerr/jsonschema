package tags

import (
	"github.com/RossMerr/jsonschema"
)

type StructTag interface {
	ToStructTag(key string, schema *jsonschema.Schema, required []string) string
}
