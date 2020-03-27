package jsonschema

import (
	"net/url"
	"path/filepath"
	"strings"
)

type Reference string

func NewReference(s string) Reference {
	return Reference(s)
}

func (s Reference) String() string {
	return string(s)
}

func (s Reference) Pointer() Pointer {
	raw := string(s)
	if len(raw) < 1 {
		return Pointer{}
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return Pointer{}
	}

	if uri.Fragment == EmptyString {
		return Pointer{}
	}
	parts := uri.Fragment

	query := strings.Split(parts, "/")
	query = Filter(query, func(v string) bool { return v != "" })
	return query
}

// Base reports the file this ID references
func (s Reference) Base() string {
	raw := string(s)
	if len(raw) < 1 {
		return "."
	}

	index := strings.Index(raw, "#")

	if index >= 0 {
		raw = raw[:index]
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return "."
	}

	file := filepath.Base(uri.Path)
	return file
}

func (s Reference) Fieldname() string {
	fragments := s.Pointer()

	if len(fragments) == 0 {
		return EmptyString
	}
	name := fragments[len(fragments)-1]

	return Title(name)
}

func (s Reference) IsNotEmpty() bool {
	return s != EmptyString
}
