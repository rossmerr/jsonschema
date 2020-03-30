package jsonschema

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MetaSchema string

const (
	// Draft 7 is not really supported
	Draft07 MetaSchema = "http://json-schema.org/draft-07/schema"
	// Draft 8 old name not really supported
	Draft08 MetaSchema = "http://json-schema.org/draft-08/schema"
	// 2019-09  formerly known as Draft 8
	IETF_2019_19 MetaSchema = "https://json-schema.org/2019-09/schema"
)

var MetaSchemas = []MetaSchema{Draft07, Draft08, IETF_2019_19}

func MetaSchemasToString() []string {
	arr := []string{}
	for _, v := range MetaSchemas {
		arr = append(arr, string(v))
	}
	return arr
}

func (s *MetaSchema) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	if !Contains(MetaSchemasToString(), v) {
		return fmt.Errorf("unsupported schema found %v\n\nTry using one of:\n%v\n", v, strings.Join(MetaSchemasToString(), ", "))
	}

	*s = MetaSchema(v)
	return nil
}
