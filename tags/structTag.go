package tags

import (
	"github.com/RossMerr/jsonschema"
)

type StructTag interface {
	ToStructTag(key string, schema, parent *jsonschema.Schema) string
}
