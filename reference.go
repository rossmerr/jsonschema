package jsonschema

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Reference ($ref) is used to reference other schemas, used with allOf, anyOf and oneOf
type Reference string

// NewReference returns a Reference
func NewReference(s string) (Reference, error) {
	return Reference(s), nil
}

func (s Reference) String() string {
	return string(s)
}

// ID return's the ID of the Reference
func (s Reference) ID() (ID, error) {
	raw := string(s)
	return NewID(raw)
}

// Path returns all segments of the relative path of the referenced schema
func (s Reference) Path() Path {
	raw := string(s)
	if len(raw) < 1 {
		return Path{}
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return Path{}
	}

	if uri.Fragment == emptyString {
		return Path{}
	}
	parts := uri.Fragment

	index := strings.Index(parts, "/")
	if index < 0 {
		return Path{}
	}
	if index != 0 {
		parts = parts[index:]
	}

	query := strings.Split(parts, "/")
	query = Filter(query, func(v string) bool { return v != "" })
	return query
}

func (s Reference) ToKey() string {
	path := s.Path()

	if len(path) == 0 {
		return emptyString
	}
	name := path[len(path)-1]

	return name
}

func (s Reference) IsNotEmpty() bool {
	return s != emptyString
}

func (s *Reference) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*s = Reference(v)
	return nil
}
