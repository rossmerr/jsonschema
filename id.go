package jsonschema

import (
	"net/url"
	"path/filepath"
	"strings"
	"unicode"
)

type ID string

func NewID(s string) ID {
	return ID(CanonicalURI(s))
}
func (s ID) String() string {
	return string(s)
}

// ToFilename returns the file name from the ID.
func (s ID) ToFilename() string {
	name := s.ToTypename()

	if len(name) > 0 {
		return string(unicode.ToLower(rune(name[0]))) + name[1:]
	}
	return name
}

func (s ID) ToTypename() string {
	raw := string(s)
	if len(raw) < 1 {
		return "."
	}

	index := strings.Index(raw, "#")
	var basename string
	if index < 0 {
		basename = filepath.Base(raw)
	} else {
		basename = filepath.Base(raw[:index])
	}


	return ToTypename(basename)
}

func CanonicalURI(s string) string {
	fragments := strings.Index(s, "#")
	if fragments > 0 {
		s = s[:fragments]
	}

	uri, err := url.Parse(s)
	if err != nil {
		return "."
	}

	if !uri.IsAbs() {
		return "."
	}

	if strings.HasSuffix(s,"/") {
		s = s[:len(s)-1]
	}

	return s
}