package jsonschema

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

type ID string

func NewID(s string) ID {
	return ID(s)
}

func (s ID) String() string {
	return string(s)
}

// Base reports the file this ID references
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

func (s ID) Filename() string {
	basename := s.Base()
	name := strings.TrimSuffix(basename, filepath.Ext(basename))

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	filename := reg.ReplaceAllString( strings.Title(clean), "")
	if len(filename) > 0 {
		return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
	}
	return filename
}
