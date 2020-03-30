package jsonschema

import (
	"encoding/json"
	"net/url"
	"strings"
)

type Reference string

func NewReference(s string) (Reference, error) {
	return Reference(s), nil
}

func (s Reference) String() string {
	return string(s)
}

func (s Reference) ID() (ID, error) {
	raw := string(s)
	return NewID(raw)
}

func (s Reference) Path() Path {
	raw := string(s)
	if len(raw) < 1 {
		return Path{}
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return Path{}
	}

	if uri.Fragment == EmptyString {
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

func (s Reference) ToTypename() string {
	path := s.Path()

	if len(path) == 0 {
		return EmptyString
	}
	name := path[len(path)-1]

	return title(name)
}

func (s Reference) IsNotEmpty() bool {
	return s != EmptyString
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
