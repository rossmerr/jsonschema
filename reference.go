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
		return Pointer("")
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return Pointer("")
	}

	if uri.Fragment == EmptyString {
		return Pointer("")
	}
	parts := uri.Fragment

	if strings.HasPrefix(parts, "/") {
		return Pointer("")
	}

	query := strings.Split(parts, "/")

	return Pointer(query[0])
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
	path := s.Path()

	if len(path) == 0 {
		return EmptyString
	}
	name := path[len(path)-1]

	return Title(name)
}

func (s Reference) IsNotEmpty() bool {
	return s != EmptyString
}

func (s Reference) Stat() (Pointer, Path) {
	return s.Pointer(), s.Path()
}
