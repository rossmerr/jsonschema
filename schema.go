package jsonschema

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/RossMerr/jsonschema/tokens"
)

type Schema struct {
	ID          ID                 `json:"$id,omitempty"`
	Schema      string             `json:"$schema,omitempty"`
	Ref         Reference          `json:"$ref,omitempty"`
	Defs        map[string]*Schema `json:"$defs,omitempty"`
	Anchor      string             `json:"$anchor,omitempty"`
	Description string             `json:"description,omitempty"`
	Title       string             `json:"title,omitempty"`
	Type        Kind               `json:"type,omitempty"`
	Required    []string           `json:"required,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	// Deprecated use Defs
	Definitions          map[string]*Schema `json:"definitions,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	OneOf                []*Schema          `json:"oneof,omitempty"`
	AnyOf                []*Schema          `json:"anyof,omitempty"`
	AllOf                []*Schema          `json:"allof,omitempty"`
	Enum                 []string           `json:"enum,omitempty"`
	AdditionalProperties *bool              `json:"additionalproperties,omitempty"`

	// Validation
	MaxProperties    *uint32 `json:"maxproperties,omitempty"`
	MinProperties    *uint32 `json:"minproperties,omitempty"`
	MaxLength        *uint32 `json:"maxlength,omitempty"`
	MinLength        *uint32 `json:"minlength,omitempty"`
	MaxContains      *uint32 `json:"maxcontains,omitempty"`
	MinContains      *uint32 `json:"mincontains,omitempty"`
	MaxItems         *uint32 `json:"maxitems,omitempty"`
	MinItems         *uint32 `json:"minitems,omitempty"`
	Maximum          *int32  `json:"maximum,omitempty"`
	ExclusiveMaximum *int32  `json:"exclusivemaximum,omitempty"`
	Minimum          *int32  `json:"minimum,omitempty"`
	ExclusiveMinimum *int32  `json:"exclusiveminimum,omitempty"`
	Pattern          string  `json:"pattern,omitempty"`
}

func (s *Schema) Stat() (Kind, Reference, []*Schema, []*Schema, []*Schema) {
	return s.Type, s.Ref, s.OneOf, s.AnyOf, s.AllOf
}

func UnmarshalSchema(data []byte) (*Schema, map[ID]*Schema, error) {
	var schema Schema
	references := map[ID]*Schema{}

	err := json.Unmarshal(data, &schema)
	if err != nil {
		return &schema, references, err
	}

	references = resolveIDs(data, []string{}, references)

	return &schema, references, err
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
