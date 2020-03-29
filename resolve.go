package jsonschema

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/RossMerr/jsonschema/tokens"
)

// ResolveIDs looks over the raw json and traverses over the schema looking for $id fields,
func ResolveIDs(b json.RawMessage) map[ID]*Schema {
	return resolveIDs(b, []string{}, map[ID]*Schema{})
}

func resolveIDs(b []byte, path []string, references map[ID]*Schema) map[ID]*Schema {
	m := map[string]json.RawMessage{}
	json.Unmarshal(b, &m)

	switch raw, ok := m[tokens.ID]; {
	case ok:
		var id string
		if err := json.Unmarshal(raw, &id); err != nil {
			break
		}

		uri, err := url.Parse(id)
		if err != nil {
			break
		}

		if uri.IsAbs() {
			path = []string{}
		}

		path = append(path, id)
		var schema Schema

		err = json.Unmarshal(b, &schema)
		if err != nil {
			break
		}

		references[ID(strings.Join(path, "/"))] = &schema
	}

	for key, raw := range m {
		if key != tokens.ID {
			for k, v := range resolveIDs(raw, append(path, key), references) {
				references[k] = v
			}
		}
	}
	return references
}
