package jsonschema

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SchemaVersion declares which version of the JSON Schema standard that the schema was written against
type SchemaVersion string

const (
	// Draft07 is not really supported
	Draft07 SchemaVersion = "http://json-schema.org/draft-07/schema"
	// Draft08 old name not really supported
	Draft08 SchemaVersion = "http://json-schema.org/draft-08/schema"
	// IETF_2019_19 formerly known as Draft 8
	IETF_2019_19 SchemaVersion = "https://json-schema.org/2019-09/schema"
)

var metaSchemas = []SchemaVersion{Draft07, Draft08, IETF_2019_19}

// MetaSchemaVersions returns a list of all supported schema versions
func MetaSchemaVersions() []string {
	arr := []string{}
	for _, v := range metaSchemas {
		arr = append(arr, string(v))
	}
	return arr
}

func (s *SchemaVersion) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	v = strings.TrimRight(v, "#")

	if !Contains(MetaSchemaVersions(), v) {
		return fmt.Errorf("unsupported schema found %v\n\nTry using one of:\n%v\n", v, strings.Join(MetaSchemaVersions(), ", "))
	}

	*s = SchemaVersion(v)
	return nil
}
