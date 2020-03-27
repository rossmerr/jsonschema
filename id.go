package jsonschema

import (
	"path/filepath"
	"strings"
	"unicode"
)

type ID string

func (s ID) String() string {
	return string(s)
}

// Base returns the last element of path.
func (s ID) Base() string {
	raw := string(s)
	if len(raw) < 1 {
		return "."
	}

	index := strings.Index(raw, "#")
	if index < 0 {
		return filepath.Base(raw)
	}

	return filepath.Base(raw[:index])
}

// Filename returns the file name from the ID.
func (s ID) Filename() string {
	basename := s.Base()

	name := Fieldname(basename)

	if len(name) > 0 {
		return string(unicode.ToLower(rune(name[0]))) + name[1:]
	}
	return name
}
